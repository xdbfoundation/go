package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/digitalbits/go/tools/digitalbits-hd-wallet/commands"
)

var mainCmd = &cobra.Command{
	Use:   "digitalbits-hd-wallet",
	Short: "Simple HD wallet for DigitalBits Lumens. THIS PROGRAM IS STILL EXPERIMENTAL. USE AT YOUR OWN RISK.",
}

func init() {
	mainCmd.AddCommand(commands.NewCmd)
	mainCmd.AddCommand(commands.AccountsCmd)
}

func main() {
	if err := mainCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
