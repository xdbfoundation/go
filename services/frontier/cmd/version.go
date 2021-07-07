package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	apkg "github.com/xdbfoundation/go/support/app"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print frontier and Golang runtime version",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(apkg.Version())
		fmt.Println(runtime.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
