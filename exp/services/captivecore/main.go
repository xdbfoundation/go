package main

import (
	"fmt"
	"go/types"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/digitalbits/go/exp/services/captivecore/internal"
	"github.com/digitalbits/go/ingest/ledgerbackend"
	"github.com/digitalbits/go/network"
	"github.com/digitalbits/go/support/config"
	"github.com/digitalbits/go/support/db"
	supporthttp "github.com/digitalbits/go/support/http"
	supportlog "github.com/digitalbits/go/support/log"
)

func main() {
	var port int
	var networkPassphrase, binaryPath, configAppendPath, dbURL string
	var historyArchiveURLs []string
	var digitalbitsCoreHTTPPort uint
	var checkpointFrequency uint32
	var logLevel logrus.Level
	logger := supportlog.New()

	configOpts := config.ConfigOptions{
		{
			Name:        "port",
			Usage:       "Port to listen and serve on",
			OptType:     types.Int,
			ConfigKey:   &port,
			FlagDefault: 8000,
			Required:    true,
		},
		{
			Name:        "network-passphrase",
			Usage:       "Network passphrase of the DigitalBits network transactions should be signed for",
			OptType:     types.String,
			ConfigKey:   &networkPassphrase,
			FlagDefault: network.TestNetworkPassphrase,
			Required:    true,
		},
		&config.ConfigOption{
			Name:        "digitalbits-core-binary-path",
			OptType:     types.String,
			FlagDefault: "",
			Required:    true,
			Usage:       "path to digitalbits core binary",
			ConfigKey:   &binaryPath,
		},
		&config.ConfigOption{
			Name:        "captive-core-config-append-path",
			OptType:     types.String,
			FlagDefault: "",
			Required:    false,
			Usage:       "path to additional configuration for the DigitalBits Core configuration file used by captive core. It must, at least, include enough details to define a quorum set",
			ConfigKey:   &configAppendPath,
		},
		&config.ConfigOption{
			Name:        "history-archive-urls",
			ConfigKey:   &historyArchiveURLs,
			OptType:     types.String,
			Required:    true,
			FlagDefault: "",
			CustomSetValue: func(co *config.ConfigOption) {
				stringOfUrls := viper.GetString(co.Name)
				urlStrings := strings.Split(stringOfUrls, ",")

				*(co.ConfigKey.(*[]string)) = urlStrings
			},
			Usage: "comma-separated list of digitalbits history archives to connect with",
		},
		&config.ConfigOption{
			Name:        "log-level",
			ConfigKey:   &logLevel,
			OptType:     types.String,
			FlagDefault: "info",
			CustomSetValue: func(co *config.ConfigOption) {
				ll, err := logrus.ParseLevel(viper.GetString(co.Name))
				if err != nil {
					logger.Fatalf("Could not parse log-level: %v", viper.GetString(co.Name))
				}
				*(co.ConfigKey.(*logrus.Level)) = ll
			},
			Usage: "minimum log severity (debug, info, warn, error) to log",
		},
		&config.ConfigOption{
			Name:      "db-url",
			EnvVar:    "DATABASE_URL",
			ConfigKey: &dbURL,
			OptType:   types.String,
			Required:  false,
			Usage:     "frontier postgres database to connect with",
		},
		&config.ConfigOption{
			Name:        "digitalbits-captive-core-http-port",
			ConfigKey:   &digitalbitsCoreHTTPPort,
			OptType:     types.Uint,
			FlagDefault: uint(11626),
			Required:    false,
			Usage:       "HTTP port for captive core to listen on (0 disables the HTTP server)",
		},
		&config.ConfigOption{
			Name:        "checkpoint-frequency",
			ConfigKey:   &checkpointFrequency,
			OptType:     types.Uint32,
			FlagDefault: uint32(64),
			Required:    false,
			Usage:       "establishes how many ledgers exist between checkpoints, do NOT change this unless you really know what you are doing",
		},
	}
	cmd := &cobra.Command{
		Use:   "captivecore",
		Short: "Run the remote captive core server",
		Run: func(_ *cobra.Command, _ []string) {
			configOpts.Require()
			configOpts.SetValues()
			logger.SetLevel(logLevel)

			captiveConfig := ledgerbackend.CaptiveCoreConfig{
				BinaryPath:          binaryPath,
				ConfigAppendPath:    configAppendPath,
				NetworkPassphrase:   networkPassphrase,
				HistoryArchiveURLs:  historyArchiveURLs,
				CheckpointFrequency: checkpointFrequency,
				HTTPPort:            digitalbitsCoreHTTPPort,
				Log:                 logger.WithField("subservice", "digitalbits-core"),
			}

			var dbConn *db.Session
			if len(dbURL) > 0 {
				var err error
				dbConn, err = db.Open("postgres", dbURL)
				if err != nil {
					logger.WithError(err).Fatal("Could not create db connection instance")
				}
				captiveConfig.LedgerHashStore = ledgerbackend.NewFrontierDBLedgerHashStore(dbConn)
			}

			core, err := ledgerbackend.NewCaptive(captiveConfig)
			if err != nil {
				logger.WithError(err).Fatal("Could not create captive core instance")
			}
			api := internal.NewCaptiveCoreAPI(core, logger.WithField("subservice", "api"))

			supporthttp.Run(supporthttp.Config{
				ListenAddr: fmt.Sprintf(":%d", port),
				Handler:    internal.Handler(api),
				OnStarting: func() {
					logger.Infof("Starting Captive Core server on %v", port)
				},
				OnStopping: func() {
					api.Shutdown()
					if dbConn != nil {
						dbConn.Close()
					}
				},
			})
		},
	}

	if err := configOpts.Init(cmd); err != nil {
		logger.WithError(err).Fatal("could not parse config options")
	}

	if err := cmd.Execute(); err != nil {
		logger.WithError(err).Fatal("could not run")
	}
}
