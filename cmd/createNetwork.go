package cmd

import (
	"github.com/spf13/cobra"
)

var createNetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "create network",
	Long:  "blah blah",
}

func init() {

	createCmd.AddCommand(createNetworkCmd)
}
