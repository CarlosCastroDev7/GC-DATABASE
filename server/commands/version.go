package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var MicroServerVersion = &cobra.Command{
	Use:   "version",
	Short: "Now the micro service version",
	Run: func(cmd *cobra.Command, args []string) {
		Name := viper.GetString("Microservice.Name")
		Version := viper.GetString("MicroService.Version")

		fmt.Printf("%s version: %s", Name, Version)
	},
}
