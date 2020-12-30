package cmd

import (
	"github.com/spf13/cobra"
	frontier "github.com/digitalbits/go/services/frontier/internal"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run frontier server",
	Long:  "serve initializes then starts the frontier HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		frontier.NewAppFromFlags(config, flags).Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
