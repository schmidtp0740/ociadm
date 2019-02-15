package cost

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/oracle/oci-go-sdk/core"

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

		compartments, err := traverseCompartments(compartmentID, []identity.Compartment{})
		if err != nil {
			log.Fatal(err.Error())
		}

		sumOfInstances := 0
		buffer := ""
		for key, comp := range compartments {
			buffer += fmt.Sprintf("%d:", key)
			buffer += fmt.Sprintf("\tName: %s\n", *comp.Name)

			// Print compartment ID to buffer
			// buffer += fmt.Sprintf("\tID: %s\n", *comp.Id)

			instancesInCompartment, err := getInstancesInCompartment(*comp.Id)
			if err != nil {
				log.Fatal(err)
			}

			sumOfInstances += len(instancesInCompartment)
			buffer += fmt.Sprintf("\t# of instances: %d\n\n", len(instancesInCompartment))
		}
		fmt.Println(buffer)
		fmt.Printf("Sum Of Instances: %d\n", sumOfInstances)

	},
}

func traverseCompartments(compartmentID string, begin []identity.Compartment) ([]identity.Compartment, error) {
	listComparmtentResponse, err := listCompartments(compartmentID)
	if err != nil {
		return nil, errors.New("Error get list of compartments: " + err.Error())
	}

	if len(listComparmtentResponse.Items) == 0 {
		return begin, nil
	}

	for _, compartment := range listComparmtentResponse.Items {
		if compartment.LifecycleState == identity.CompartmentLifecycleStateActive {
			begin = append(begin, compartment)
			begin, err = traverseCompartments(*compartment.Id, begin)
			if err != nil {
				return nil, err
			}
		}
	}

	return begin, nil
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

func getInstancesInCompartment(compID string) ([]core.Instance, error) {
	client, err := pkg.GetDefaultComputeClient()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	request := core.ListInstancesRequest{
		CompartmentId: common.String(compID),
	}

	response, err := client.ListInstances(ctx, request)
	if err != nil {
		return nil, err
	}

	return response.Items, nil

}
