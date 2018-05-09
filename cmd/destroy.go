package cmd

import (
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "TODO",
	Long:  "TODO",
}

func init() {
	RootCmd.AddCommand(destroyCmd)
}
