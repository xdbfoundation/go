//lint:file-ignore U1001 Ignore all unused code, this is only used in tests.
package integration

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	sdk "github.com/xdbfoundation/go/clients/frontierclient"
	"github.com/xdbfoundation/go/clients/digitalbitscore"
	"github.com/xdbfoundation/go/keypair"
	proto "github.com/xdbfoundation/go/protocols/frontier"
	frontier "github.com/xdbfoundation/go/services/frontier/internal"
	"github.com/xdbfoundation/go/support/db/dbtest"
	"github.com/xdbfoundation/go/txnbuild"
	"github.com/xdbfoundation/go/xdr"
)

const (
	NetworkPassphrase           = "Standalone Network ; February 2017"
	digitalbitsCorePostgresPassword = "mysecretpassword"
	adminPort                   = 6060
	digitalbitsCorePort             = 11626
	digitalbitsCorePostgresPort     = 5641
	historyArchivePort          = 1570
)

type Config struct {
	PostgresURL           string
	ProtocolVersion       int32
	SkipContainerCreation bool
}

type Test struct {
	t             *testing.T
	config        Config
	frontierConfig frontier.Config
	hclient       *sdk.Client
	cclient       *digitalbitscore.Client
	app           *frontier.App
	appStopped    chan struct{}
	shutdownOnce  sync.Once
	shutdownCalls []func()
}

// NewTest starts a new environment for integration test at a given
// protocol version and blocks until Frontier starts ingesting.
//
// Warning: this requires Docker Compose installed
//
// Skips the test if FRONTIER_INTEGRATION_TESTS env variable is not set.
func NewTest(t *testing.T, config Config) *Test {
	if os.Getenv("FRONTIER_INTEGRATION_TESTS") == "" {
		t.Skip("skipping integration test")
	}

	i := &Test{t: t, config: config}

	composeDir := findDockerComposePath()
	integrationYaml := filepath.Join(composeDir, "docker-compose.integration-tests.yml")

	// Runs a docker-compose command applied to the above configs
	runComposeCommand := func(args ...string) {
		cmdline := append([]string{"-f", integrationYaml}, args...)
		t.Log("Running", cmdline)
		cmd := exec.Command("docker-compose", cmdline...)
		_, innerErr := cmd.Output()
		fatalIf(t, innerErr)
	}

	var captiveCoreBinaryPath, captiveCoreConfigPath string
	if os.Getenv("FRONTIER_INTEGRATION_ENABLE_CAPTIVE_CORE") != "" {
		captiveCoreBinaryPath = os.Getenv("CAPTIVE_CORE_BIN")
		if len(captiveCoreBinaryPath) == 0 {
			t.Fatal("CAPTIVE_CORE_BIN is not set")
		}
		captiveCoreConfigPath = filepath.Join(composeDir, "captive-core-integration-tests.cfg")
	}

	// Only run DigitalBits Core container and its dependencies
	runComposeCommand("up", "--detach", "--quiet-pull", "--no-color", "core")

	i.cclient = &digitalbitscore.Client{URL: "http://localhost:" + strconv.Itoa(digitalbitsCorePort)}
	i.waitForCore()

	i.shutdownCalls = append(i.shutdownCalls, func() {
		if i.app != nil {
			i.app.Close()
		}
		runComposeCommand("down", "-v", "--remove-orphans")
	})
	i.startFrontier(
		captiveCoreBinaryPath,
		captiveCoreConfigPath,
		i.config.PostgresURL,
		true,
	)
	i.hclient = &sdk.Client{FrontierURL: "http://localhost:8000"}

	// Register cleanup handlers (on panic and ctrl+c) so the containers are
	// stopped even if ingestion or testing fails.
	i.t.Cleanup(i.Shutdown)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		i.Shutdown()
		os.Exit(int(syscall.SIGTERM))
	}()

	i.waitForFrontier()
	return i
}

func (i *Test) RestartFrontier() {
	i.app.Close()

	// wait for frontier to shut down completely
	<-i.appStopped

	i.startFrontier(
		i.frontierConfig.CaptiveCoreBinaryPath,
		i.frontierConfig.CaptiveCoreConfigAppendPath,
		i.frontierConfig.DatabaseURL,
		false,
	)
	i.waitForFrontier()
}

func (i *Test) GetFrontierConfig() frontier.Config {
	return i.frontierConfig
}

// Shutdown stops the integration tests and destroys all its associated resources.
// Shutdown() will be implicitly called when the calling test (i.e. the `testing.Test` passed
// to `New()`) is finished if it hasn't been explicitly called before.
func (i *Test) Shutdown() {
	i.shutdownOnce.Do(func() {
		// run them in the opposite order in which they where added
		for callI := len(i.shutdownCalls) - 1; callI >= 0; callI-- {
			i.shutdownCalls[callI]()
		}
	})
}

func (i *Test) startFrontier(
	captiveCoreBinaryPath, captiveCoreConfigPath, frontierPostgresURL string, buildGenesisState bool,
) {
	if frontierPostgresURL == "" {
		postgres := dbtest.Postgres(i.t)
		i.shutdownCalls = append(i.shutdownCalls, func() {
			// TODO: Unfortunately Frontier leaves open sessions behind leading to
			//       a "database  is being accessed by other users"
			//       error when trying to drop it
			// postgres.Close()
		})
		frontierPostgresURL = postgres.DSN
	}

	config, configOpts := frontier.Flags()
	cmd := &cobra.Command{
		Use:   "frontier",
		Short: "client-facing api server for the digitalbits network",
		Long: `client-facing api server for the digitalbits network. It acts as the
interface between DigitalBits Core and applications that want to access the DigitalBits
network. It allows you to submit transactions to the network, check the status
of accounts, subscribe to event streams and more.`,
		Run: func(cmd *cobra.Command, args []string) {
			i.app = frontier.NewAppFromFlags(config, configOpts)
		},
	}

	// Ideally, we'd be pulling host/port information from the Docker Compose
	// YAML file itself rather than hardcoding it.
	hostname := "localhost"
	args := []string{
		"--digitalbits-core-url",
		fmt.Sprintf("http://%s:%d", hostname, digitalbitsCorePort),
		"--history-archive-urls",
		fmt.Sprintf("http://%s:%d", hostname, historyArchivePort),
		"--ingest",
		"--db-url",
		frontierPostgresURL,
		"--digitalbits-core-db-url",
		fmt.Sprintf(
			"postgres://postgres:%s@%s:%d/digitalbits?sslmode=disable",
			digitalbitsCorePostgresPassword,
			hostname,
			digitalbitsCorePostgresPort,
		),

		"--digitalbits-core-binary-path",
		captiveCoreBinaryPath,
		"--captive-core-config-append-path",
		captiveCoreConfigPath,

		// disable http port to not clash with the http port of the
		// non-captive digitalbits core instance running in docker
		"--captive-core-http-port",
		"0",

		"--enable-captive-core-ingestion=" + strconv.FormatBool(len(captiveCoreBinaryPath) > 0),

		"--network-passphrase",
		NetworkPassphrase,
		"--apply-migrations",
		"--admin-port",
		strconv.Itoa(i.AdminPort()),

		// due to ARTIFICIALLY_ACCELERATE_TIME_FOR_TESTING
		"--checkpoint-frequency",
		"8",
	}

	// initialize core arguments
	cmd.SetArgs(args)
	var err error
	if err = configOpts.Init(cmd); err != nil {
		i.t.Fatalf("Cannot initialize params: %s", err)
	}

	if err = cmd.Execute(); err != nil {
		i.t.Fatalf("cannot initialize frontier: %s", err)
	}
	i.frontierConfig = *config

	if buildGenesisState {
		if err = i.app.Ingestion().BuildGenesisState(); err != nil {
			i.t.Fatalf("cannot build genesis state: %s", err)
		}
	}

	done := make(chan struct{})
	go func() {
		i.app.Serve()
		close(done)
	}()
	i.appStopped = done
}

// Wait for core to be up and manually close the first ledger
func (i *Test) waitForCore() {
	for t := 30 * time.Second; t >= 0; t -= time.Second {
		i.t.Log("Waiting for core to be up...")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		_, err := i.cclient.Info(ctx)
		cancel()
		if err != nil {
			i.t.Logf("could not obtain info response: %v", err)
			time.Sleep(time.Second)
			continue
		}
		break
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		err := i.cclient.Upgrade(ctx, int(i.config.ProtocolVersion))
		cancel()
		if err != nil {
			i.t.Fatalf("could not upgrade protocol: %v", err)
		}
	}

	for t := 0; t < 5; t++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		info, err := i.cclient.Info(ctx)
		cancel()
		if err != nil || !info.IsSynced() {
			i.t.Logf("Core is still not synced: %v %v", err, info)
			time.Sleep(time.Second)
			continue
		}
		return
	}
	i.t.Fatal("Core could not sync after several attempts")
}

func (i *Test) waitForFrontier() {
	for t := 200; t >= 0; t -= 1 {
		i.t.Log("Waiting for ingestion and protocol upgrade...")
		root, err := i.hclient.Root()
		if err != nil {
			i.t.Logf("could not obtain root response %v", err)
			time.Sleep(time.Second)
			continue
		}

		if root.FrontierSequence < 3 ||
			int(root.FrontierSequence) != int(root.IngestSequence) {
			i.t.Logf("Frontier ingesting... %v", root)
			time.Sleep(time.Second)
			continue
		}

		if root.CurrentProtocolVersion == i.config.ProtocolVersion {
			i.t.Logf("Frontier protocol version matches... %v", root)
			return
		}
	}

	i.t.Fatal("Frontier not ingesting...")
}

// Client returns frontier.Client connected to started Frontier instance.
func (i *Test) Client() *sdk.Client {
	return i.hclient
}

// Frontier returns the frontier.App instance for the current integration test
func (i *Test) Frontier() *frontier.App {
	return i.app
}

// AdminPort returns Frontier admin port.
func (i *Test) AdminPort() int {
	return adminPort
}

// Metrics URL returns Frontier metrics URL.
func (i *Test) MetricsURL() string {
	return fmt.Sprintf("http://localhost:%d/metrics", i.AdminPort())
}

// Master returns a keypair of the network master account.
func (i *Test) Master() *keypair.Full {
	return keypair.Master(NetworkPassphrase).(*keypair.Full)
}

func (i *Test) MasterAccount() txnbuild.Account {
	master, client := i.Master(), i.Client()
	request := sdk.AccountRequest{AccountID: master.Address()}
	account, err := client.AccountDetail(request)
	panicIf(err)
	return &account
}

func (i *Test) CurrentTest() *testing.T {
	return i.t
}

/* Utility functions for easier test case creation. */

// Creates new accounts via the master account.
//
// It funds each account with the given balance and then queries the API to
// find the randomized sequence number for future operations.
//
// Returns: The slice of created keypairs and account objects.
//
// Note: panics on any errors, since we assume that tests cannot proceed without
// this method succeeding.
func (i *Test) CreateAccounts(count int, initialBalance string) ([]*keypair.Full, []txnbuild.Account) {
	client := i.Client()
	master := i.Master()

	pairs := make([]*keypair.Full, count)
	ops := make([]txnbuild.Operation, count)

	// Two paths here: either caller already did some stuff with the master
	// account so we should retrieve the sequence number, or caller hasn't and
	// we start from scratch.
	seq := int64(0)
	request := sdk.AccountRequest{AccountID: master.Address()}
	account, err := client.AccountDetail(request)
	if err == nil {
		seq, err = strconv.ParseInt(account.Sequence, 10, 8) // why is this a string?
		panicIf(err)
	}

	masterAccount := txnbuild.SimpleAccount{
		AccountID: master.Address(),
		Sequence:  seq,
	}

	for i := 0; i < count; i++ {
		pair, _ := keypair.Random()
		pairs[i] = pair

		ops[i] = &txnbuild.CreateAccount{
			SourceAccount: &masterAccount,
			Destination:   pair.Address(),
			Amount:        initialBalance,
		}
	}

	// Submit transaction, then retrieve new account details.
	_ = i.MustSubmitOperations(&masterAccount, master, ops...)

	accounts := make([]txnbuild.Account, count)
	for i, kp := range pairs {
		request := sdk.AccountRequest{AccountID: kp.Address()}
		account, err := client.AccountDetail(request)
		panicIf(err)

		accounts[i] = &account
	}

	for _, keys := range pairs {
		i.t.Logf("Funded %s (%s) with %s XDB.\n",
			keys.Seed(), keys.Address(), initialBalance)
	}

	return pairs, accounts
}

// Panics on any error establishing a trustline.
func (i *Test) MustEstablishTrustline(
	truster *keypair.Full, account txnbuild.Account, asset txnbuild.Asset,
) (resp proto.Transaction) {
	txResp, err := i.EstablishTrustline(truster, account, asset)
	panicIf(err)
	return txResp
}

// Establishes a trustline for a given asset for a particular account.
func (i *Test) EstablishTrustline(
	truster *keypair.Full, account txnbuild.Account, asset txnbuild.Asset,
) (proto.Transaction, error) {
	if asset.IsNative() {
		return proto.Transaction{}, nil
	}
	return i.SubmitOperations(account, truster, &txnbuild.ChangeTrust{
		Line:  asset,
		Limit: "2000",
	})
}

// Panics on any error creating a claimable balance.
func (i *Test) MustCreateClaimableBalance(
	source *keypair.Full, asset txnbuild.Asset, amount string,
	claimants ...txnbuild.Claimant,
) (claim proto.ClaimableBalance) {
	account := i.MustGetAccount(source)
	_ = i.MustSubmitOperations(&account, source,
		&txnbuild.CreateClaimableBalance{
			Destinations: claimants,
			Asset:        asset,
			Amount:       amount,
		},
	)

	// Ensure it exists in the global list
	balances, err := i.Client().ClaimableBalances(sdk.ClaimableBalanceRequest{})
	panicIf(err)

	claims := balances.Embedded.Records
	if len(claims) == 0 {
		panic(-1)
	}

	claim = claims[len(claims)-1] // latest one
	i.t.Logf("Created claimable balance w/ id=%s", claim.BalanceID)
	return
}

// Panics on any error retrieves an account's details from its key.
// This means it must have previously been funded.
func (i *Test) MustGetAccount(source *keypair.Full) proto.Account {
	client := i.Client()
	account, err := client.AccountDetail(sdk.AccountRequest{AccountID: source.Address()})
	panicIf(err)
	return account
}

// Submits a signed transaction from an account with standard options.
//
// Namely, we set the standard fee, time bounds, etc. to "non-production"
// defaults that work well for tests.
//
// Most transactions only need one signer, so see the more verbose
// `MustSubmitOperationsWithSigners` below for multi-sig transactions.
//
// Note: We assume that transaction will be successful here so we panic in case
// of all errors. To allow failures, use `SubmitOperations`.
func (i *Test) MustSubmitOperations(
	source txnbuild.Account, signer *keypair.Full, ops ...txnbuild.Operation,
) proto.Transaction {
	tx, err := i.SubmitOperations(source, signer, ops...)
	panicIf(err)
	return tx
}

func (i *Test) SubmitOperations(
	source txnbuild.Account, signer *keypair.Full, ops ...txnbuild.Operation,
) (proto.Transaction, error) {
	return i.SubmitMultiSigOperations(source, []*keypair.Full{signer}, ops...)
}

func (i *Test) SubmitMultiSigOperations(
	source txnbuild.Account, signers []*keypair.Full, ops ...txnbuild.Operation,
) (proto.Transaction, error) {
	tx, err := i.CreateSignedTransaction(source, signers, ops...)
	if err != nil {
		return proto.Transaction{}, err
	}
	return i.Client().SubmitTransaction(tx)
}

func (i *Test) CreateSignedTransaction(
	source txnbuild.Account, signers []*keypair.Full, ops ...txnbuild.Operation,
) (*txnbuild.Transaction, error) {
	txParams := txnbuild.TransactionParams{
		SourceAccount:        source,
		Operations:           ops,
		BaseFee:              txnbuild.MinBaseFee,
		Timebounds:           txnbuild.NewInfiniteTimeout(),
		IncrementSequenceNum: true,
	}

	tx, err := txnbuild.NewTransaction(txParams)
	if err != nil {
		return nil, err
	}

	for _, signer := range signers {
		tx, err = tx.Sign(NetworkPassphrase, signer)
		if err != nil {
			return nil, err
		}
	}

	return tx, nil
}

func (i *Test) GetCurrentCoreLedgerSequence() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	info, err := i.cclient.Info(ctx)
	if err != nil {
		return 0, err
	}
	return info.Info.Ledger.Num, nil
}

// A convenience function to provide verbose information about a failing
// transaction to the test output log, if it's expected to succeed.
func (i *Test) LogFailedTx(txResponse proto.Transaction, frontierResult error) {
	t := i.CurrentTest()
	assert.NoErrorf(t, frontierResult, "Submitting the transaction failed")
	if prob := sdk.GetError(frontierResult); prob != nil {
		t.Logf("  problem: %s\n", prob.Problem.Detail)
		t.Logf("  extras: %s\n", prob.Problem.Extras["result_codes"])
		return
	}

	var txResult xdr.TransactionResult
	err := xdr.SafeUnmarshalBase64(txResponse.ResultXdr, &txResult)
	assert.NoErrorf(t, err, "Unmarshalling transaction failed.")
	assert.Equalf(t, xdr.TransactionResultCodeTxSuccess, txResult.Result.Code,
		"Transaction doesn't have success code.")
}

// Cluttering code with if err != nil is absolute nonsense.
func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func fatalIf(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("error: %s", err)
	}
}

// Performs a best-effort attempt to find the project's Docker Compose files.
func findDockerComposePath() string {
	// Lets you check if a particular directory contains a file.
	directoryContainsFilename := func(dir string, filename string) bool {
		files, innerErr := ioutil.ReadDir(dir)
		panicIf(innerErr)

		for _, file := range files {
			if file.Name() == filename {
				return true
			}
		}

		return false
	}

	current, err := os.Getwd()
	panicIf(err)

	//
	// We have a primary and backup attempt for finding the necessary docker
	// files: via $GOPATH and via local directory traversal.
	//

	if gopath := os.Getenv("GOPATH"); gopath != "" {
		monorepo := filepath.Join(gopath, "digitalbits", "go")
		if _, err = os.Stat(monorepo); !os.IsNotExist(err) {
			current = monorepo
		}
	}

	// In either case, we try to walk up the tree until we find "go.mod",
	// which we hope is the root directory of the project.
	for !directoryContainsFilename(current, "go.mod") {
		current, err = filepath.Abs(filepath.Join(current, ".."))

		// FIXME: This only works on *nix-like systems.
		if err != nil || filepath.Base(current)[0] == filepath.Separator {
			fmt.Println("Failed to establish project root directory.")
			panic(err)
		}
	}

	// Directly jump down to the folder that should contain the configs
	return filepath.Join(current, "services", "frontier", "docker")
}
