package network

import (
	"github.com/schmidtp0740/ociadm/cmd/services/network/vcn"
	"github.com/spf13/cobra"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "network",
	Short: "TODO",
	Long:  "TODO",
}

func init() {
	Cmd.AddCommand(vcn.Cmd)
}
