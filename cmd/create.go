package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create services(",
	Long:  `This subcommand allows you to create services`,
}

func init() {

	RootCmd.AddCommand(createCmd)
}
