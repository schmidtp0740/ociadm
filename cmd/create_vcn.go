package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/spf13/cobra"
)

func init() {
	var cidrBlock, displayName, dnsLabel, compartmentId string
	vcn := core.CreateVcnDetails{
		CidrBlock:     &cidrBlock,
		DisplayName:   &displayName,
		DnsLabel:      &dnsLabel,
		CompartmentId: &compartmentId,
	}

	var createVcnCmd = &cobra.Command{
		Use:   "vcn",
		Short: "TODO",
		Long:  "TODO",

		Run: func(cmd *cobra.Command, args []string) {

			vcnDetails := createVCN(vcn)
			fmt.Printf("VCN OCID: %s\n", *vcnDetails.Id)
			fmt.Printf("VCN DisplayName: %s\n", *vcnDetails.DisplayName)
			fmt.Printf("VCN CIDR Block: %s\n", *vcnDetails.CidrBlock)
			fmt.Printf("VCN Lifecycle: %s\n", vcnDetails.LifecycleState)

		},
	}
	createVcnCmd.Flags().StringVarP(&cidrBlock, "cidr-block", "c", "10.0.0.0/16", "cidrBlock")
	createVcnCmd.Flags().StringVarP(&displayName, "name", "n", "vcn", "VCN Display Name")
	createVcnCmd.Flags().StringVarP(&dnsLabel, "dns", "d", "dnsLabel", "dnsLabel")
	createVcnCmd.Flags().StringVarP(&compartmentId, "compartment-id", "i", os.Getenv("TF_VAR_compartment_ocid"), "ocid of compartment")
	CreateCmd.AddCommand(createVcnCmd)
}

func createVCN(vcn core.CreateVcnDetails) core.Vcn {
	configProvider := common.ConfigurationProviderEnvironmentVariables("TF_VAR", "")

	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(configProvider)
	if err != nil {
		fmt.Print("Error: ")
		panic(err)
	}

	ctx := context.Background()

	request := core.CreateVcnRequest{}

	request.CreateVcnDetails = vcn

	response, err := client.CreateVcn(ctx, request)
	if err != nil {
		fmt.Println("error at create vcn")
		panic(err)
	}

	return response.Vcn
}
