package cost

import (
	"context"
	"errors"
	"fmt"

	"github.com/schmidtp0740/ociadm/pkg"
	"github.com/spf13/cobra"
)

// Cmd ...
// cobra command details to get report of usage in compartment
var Cmd = &cobra.Command{
	Use:  "cost [compartment-id]",
	Long: "retrieve a report of used services in OCI within the compartment given and its child compartments",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		compartmentID := args[0]
		fmt.Printf("compartment id used: %s\n", compartmentID)

		listCompartments(compartmentID)
	},
}

// getCompartments ...
// using parent compartment ocid retreive and
// traverse child compartments
func listCompartments(compartmentID string) error {
	// get client
	client, err := pkg.GetDefaultIdentityClient()
	if err != nil {
		return errors.New("Not able to authenticate: " + err.Error())
	}

	ctx := context.Background()

	// get list of compartments using provided ocid

	return nil

}
