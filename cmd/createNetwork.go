package cmd

import (
	"github.com/spf13/cobra"
)

var createNetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "TODO",
	Long:  "TODO",
}

func init() {

	createCmd.AddCommand(createNetworkCmd)
}
