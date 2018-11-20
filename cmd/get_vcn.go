package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/spf13/cobra"
)

func init() {
	var compartmentId, displayName, lifeCycleState string

	var getVcnCmd = &cobra.Command{
		Use:   "vcn",
		Short: "TODO",
		Long:  "TODO",

		Run: func(cmd *cobra.Command, args []string) {

			listOfVcn := getVcn(compartmentId, displayName, lifeCycleState)

			fmt.Printf("%-30s%-20s%-80s\n", "Display Name", "CIDR Block", "OCID")
			for _, item := range listOfVcn {
				fmt.Printf("%-30s%-20s%-80s\n", *item.DisplayName, *item.CidrBlock, *item.Id)
			}

		},
	}
	getVcnCmd.Flags().StringVarP(&displayName, "name", "n", "", "name of the vcn")
	getVcnCmd.Flags().StringVarP(&lifeCycleState, "lifecycle-state", "l", "", "lifecycle state")
	getVcnCmd.Flags().StringVarP(&compartmentId, "compartment-id", "i", os.Getenv("TF_VAR_compartment_ocid"), "ocid of compartment")
	GetCmd.AddCommand(getVcnCmd)
}

func getVcn(compartmentId, displayName, lifeCycleState string) []core.Vcn {
	configProvider := common.ConfigurationProviderEnvironmentVariables("TF_VAR", "")

	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(configProvider)
	if err != nil {
		fmt.Print("Error: ")
		panic(err)
	}

	ctx := context.Background()

	request := core.ListVcnsRequest{}

	request.CompartmentId = &compartmentId

	if displayName != "" {
		request.DisplayName = &displayName
	}

	if strings.ToLower(lifeCycleState) == "available" {
		request.LifecycleState = core.VcnLifecycleStateAvailable
	} else if strings.ToLower(lifeCycleState) == "provisioning" {
		request.LifecycleState = core.VcnLifecycleStateProvisioning
	} else if strings.ToLower(lifeCycleState) == "terminated" {
		request.LifecycleState = core.VcnLifecycleStateTerminated
	} else if strings.ToLower(lifeCycleState) == "terminating" {
		request.LifecycleState = core.VcnLifecycleStateTerminating
	}

	response, err := client.ListVcns(ctx, request)
	if err != nil {
		fmt.Println("error at create vcn")
		panic(err)
	}

	return response.Items
}
