package integration

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/digitalbits/go/historyarchive"
	frontiercmd "github.com/digitalbits/go/services/frontier/cmd"
	"github.com/digitalbits/go/services/frontier/internal/db2/schema"
	"github.com/digitalbits/go/services/frontier/internal/test/integration"
	"github.com/digitalbits/go/support/db"
	"github.com/digitalbits/go/support/db/dbtest"
	"github.com/digitalbits/go/txnbuild"
)

func initializeDBIntegrationTest(t *testing.T) (itest *integration.Test, reachedLedger int32) {
	itest = integration.NewTest(t, protocol15Config)
	master := itest.Master()
	tt := assert.New(t)

	// Initialize the database with some ledgers including some transactions we submit
	op := txnbuild.Payment{
		Destination: master.Address(),
		Amount:      "10",
		Asset:       txnbuild.NativeAsset{},
	}
	// TODO: should we enforce certain number of ledgers to be ingested?
	for i := 0; i < 8; i++ {
		txResp := itest.MustSubmitOperations(itest.MasterAccount(), master, &op)
		reachedLedger = txResp.Ledger
	}

	root, err := itest.Client().Root()
	tt.NoError(err)
	tt.LessOrEqual(reachedLedger, root.FrontierSequence)

	return
}

func TestReingestDB(t *testing.T) {
	itest, reachedLedger := initializeDBIntegrationTest(t)
	tt := assert.New(t)

	// Create a fresh Frontier database
	newDB := dbtest.Postgres(t)
	// TODO: Unfortunately Frontier's ingestion System leaves open sessions behind,leading to
	//       a "database  is being accessed by other users" error when trying to drop it
	// defer newDB.Close()
	freshFrontierPostgresURL := newDB.DSN
	frontierConfig := itest.GetFrontierConfig()
	frontierConfig.DatabaseURL = freshFrontierPostgresURL
	// Initialize the DB schema
	dbConn, err := db.Open("postgres", freshFrontierPostgresURL)
	defer dbConn.Close()
	_, err = schema.Migrate(dbConn.DB.DB, schema.MigrateUp, 0)
	tt.NoError(err)

	// cap reachedLedger to the nearest checkpoint ledger because reingest range cannot ingest past the most
	// recent checkpoint ledger when using captive core
	toLedger := uint32(reachedLedger)
	archive, err := historyarchive.Connect(frontierConfig.HistoryArchiveURLs[0], historyarchive.ConnectOptions{
		NetworkPassphrase:   frontierConfig.NetworkPassphrase,
		CheckpointFrequency: frontierConfig.CheckpointFrequency,
	})
	tt.NoError(err)

	// make sure a full checkpoint has elapsed otherwise there will be nothing to reingest
	var latestCheckpoint uint32
	publishedFirstCheckpoint := func() bool {
		has, requestErr := archive.GetRootHAS()
		tt.NoError(requestErr)
		latestCheckpoint = has.CurrentLedger
		return latestCheckpoint > 1
	}
	tt.Eventually(publishedFirstCheckpoint, 10*time.Second, time.Second)

	if toLedger > latestCheckpoint {
		toLedger = latestCheckpoint
	}

	frontierConfig.CaptiveCoreConfigAppendPath = filepath.Join(
		filepath.Dir(frontierConfig.CaptiveCoreConfigAppendPath),
		"captive-core-reingest-range-integration-tests.cfg",
	)

	// Reingest into the DB
	err = frontiercmd.RunDBReingestRange(1, toLedger, false, 1, frontierConfig)
	tt.NoError(err)
}

func TestResumeFromInitializedDB(t *testing.T) {
	itest, reachedLedger := initializeDBIntegrationTest(t)
	tt := assert.New(t)

	// Stop the integration test, and restart it with the same database
	oldDBURL := itest.GetFrontierConfig().DatabaseURL
	itestConfig := protocol15Config
	itestConfig.PostgresURL = oldDBURL

	itest.RestartFrontier()

	successfullyResumed := func() bool {
		root, err := itest.Client().Root()
		tt.NoError(err)
		// It must be able to reach the ledger and surpass it
		const ledgersPastStopPoint = 4
		return root.FrontierSequence > (reachedLedger + ledgersPastStopPoint)
	}

	tt.Eventually(successfullyResumed, 1*time.Minute, 1*time.Second)
}
