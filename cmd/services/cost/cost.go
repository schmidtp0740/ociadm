package cost

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/identity"
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

		listcomparmtentResponse, err := listCompartments(compartmentID)
		if err != nil {
			log.Fatal("Error get list of compartments: " + err.Error())
		}

		for key, compartment := range listcomparmtentResponse.Items {
			if compartment.LifecycleState == identity.CompartmentLifecycleStateActive {
				fmt.Printf("%d : %s\n", key, *compartment.Name)

			}
		}
	},
}

// getCompartments ...
// using parent compartment ocid retreive and
// traverse child compartments
func listCompartments(compartmentID string) (identity.ListCompartmentsResponse, error) {
	// get client
	client, err := pkg.GetDefaultIdentityClient()
	if err != nil {
		return identity.ListCompartmentsResponse{}, errors.New("Not able to authenticate: " + err.Error())
	}

	ctx := context.Background()

	request := identity.ListCompartmentsRequest{}

	request.CompartmentId = common.String(compartmentID)

	// get list of compartments using provided ocid

	response, err := client.ListCompartments(ctx, request)
	if err != nil {
		return identity.ListCompartmentsResponse{}, err
	}

	return response, nil

}
