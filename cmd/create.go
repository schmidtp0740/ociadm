package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "TODO",
	Long:  "TODO",
}

func init() {

	RootCmd.AddCommand(createCmd)
}
