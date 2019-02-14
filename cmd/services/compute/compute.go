package compute

import (
	"github.com/schmidtp0740/ociadm/cmd/services/compute/instance"
	"github.com/spf13/cobra"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "compute",
	Short: "TODO",
	Long:  "TODO",
}

func init() {
	Cmd.AddCommand(instance.Cmd)
}
