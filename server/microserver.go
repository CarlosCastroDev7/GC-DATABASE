package server

import (
	"fmt"
	"gc/database/server/commands"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "all",
	Short: "All Command Application",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Start() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(commands.MicroServerVersion)
	rootCmd.AddCommand(commands.MicroServerApi)
}
