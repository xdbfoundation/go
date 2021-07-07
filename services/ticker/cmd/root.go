package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	frontierclient "github.com/xdbfoundation/go/clients/frontierclient"
	hlog "github.com/xdbfoundation/go/support/log"
)

var DatabaseURL string
var Client *frontierclient.Client
var UseTestNet bool
var Logger = hlog.New()

var defaultDatabaseURL = getEnv("DB_URL", "postgres://localhost:5432/digitalbitsticker01?sslmode=disable")

var rootCmd = &cobra.Command{
	Use:   "ticker",
	Short: "DigitalBits Development Foundation Ticker.",
	Long:  `A tool to provide DigitalBits Asset and Market data.`,
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(
		&DatabaseURL,
		"db-url",
		"d",
		defaultDatabaseURL,
		"database URL, such as: postgres://user:pass@localhost:5432/ticker",
	)
	rootCmd.PersistentFlags().BoolVar(
		&UseTestNet,
		"testnet",
		false,
		"use the DigitalBits Test Network, instead of the DigitalBits Public Network",
	)

	Logger.SetLevel(logrus.DebugLevel)
}

func initConfig() {
	if UseTestNet {
		Logger.Debug("Using DigitalBits Default Test Network")
		Client = frontierclient.DefaultTestNetClient
	} else {
		Logger.Debug("Using DigitalBits Default Public Network")
		Client = frontierclient.DefaultPublicNetClient
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
