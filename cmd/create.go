package cmd

import (
	"github.com/spf13/cobra"
)

// CreateCmd ...
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "TODO",
	Long:  "TODO",
}

func init() {

	RootCmd.AddCommand(CreateCmd)
}
