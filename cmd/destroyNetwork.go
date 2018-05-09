package cmd

import "github.com/spf13/cobra"

var destroyNetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "TODO",
	Long:  "TODO",
}

func init() {
	destroyCmd.AddCommand(destroyNetworkCmd)
}
