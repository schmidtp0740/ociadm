package cmd

import (
	"github.com/spf13/cobra"
)

// GetCmd ...
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "TODO",
	Long:  "TODO",
}

func init() {

	RootCmd.AddCommand(GetCmd)
}
