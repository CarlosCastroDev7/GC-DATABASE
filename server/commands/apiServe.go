package commands

import (
	"gc/database/api"

	"github.com/spf13/cobra"
)

var MicroServerApi = &cobra.Command{
	Use:   "start",
	Short: "Start the micro service api",
	Run: func(cmd *cobra.Command, args []string) {
		api.StartAPI()
	},
}
