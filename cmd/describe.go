package cmd

import (
	"github.com/spf13/cobra"
)

// DestroyCmd ...
var DescribeCmd = &cobra.Command{
	Use:   "describe",
	Short: "TODO",
	Long:  "TODO",
}

func init() {
	RootCmd.AddCommand(DescribeCmd)
}
