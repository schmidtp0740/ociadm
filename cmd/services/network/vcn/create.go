package vcn

import (
	"context"
	"errors"
	"fmt"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/spf13/cobra"
)

// CreateCmd cobra command for creating a vcn
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "TODO",
	Long:  "TODO",

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("test1")

		// request := core.CreateVcnRequest{}

		// cidrBlock, err := cmd.Flags().GetString("cidr-block")
		// if err != nil {
		// 	fmt.Print("Error: Unable to get cidr-block flag")
		// 	return
		// }
		// request.CidrBlock = common.String(cidrBlock)

		// displayName, err := cmd.Flags().GetString("name")
		// if err != nil {
		// 	fmt.Print("Error: Unable to get display-name")
		// 	return
		// }
		// request.DisplayName = common.String(displayName)

		// dnsLabel, err := cmd.Flags().GetString("dns")
		// if err != nil {
		// 	fmt.Print("Error: Unable to get dns-label")
		// 	return
		// }
		// request.DnsLabel = common.String(dnsLabel)

		// compartmentID, err := cmd.Flags().GetString("compartment-id")
		// if err != nil {
		// 	fmt.Print("Error: Unable to get compartment-id")
		// 	return
		// }
		// request.CompartmentId = common.String(compartmentID)

		// response, err := CreateVcn(request)
		// if err != nil {
		// 	fmt.Printf(err.Error())
		// }

		// fmt.Printf("VCN OCID: %s\n", *response.Id)
		// fmt.Printf("VCN DisplayName: %s\n", *response.DisplayName)
		// fmt.Printf("VCN CIDR Block: %s\n", *response.CidrBlock)
		// fmt.Printf("VCN Lifecycle: %s\n", response.LifecycleState)

	},
}

func init() {
	var cidrBlock, displayName, dnsLabel, compartmentID string
	CreateCmd.Flags().StringVarP(&cidrBlock, "cidr-block", "c", "10.0.0.0/16", "cidrBlock")
	CreateCmd.Flags().StringVarP(&displayName, "name", "n", "defaultvcn", "VCN Display Name")
	CreateCmd.Flags().StringVarP(&dnsLabel, "dns", "d", "vcn", "dnsLabel")
	CreateCmd.Flags().StringVarP(&compartmentID, "compartment-id", "i", "1", "ocid of compartment")
	Cmd.AddCommand(CreateCmd)
}

// CreateVcn creates a vcn
// Input: details for creating the vcn
// Output: result from oci control plane
func CreateVcn(request core.CreateVcnRequest) (core.CreateVcnResponse, error) {
	configProvider := common.DefaultConfigProvider()

	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(configProvider)
	if err != nil {
		return core.CreateVcnResponse{}, errors.New("Unable to get configuration provider" + err.Error())
	}

	ctx := context.Background()

	response, err := client.CreateVcn(ctx, request)
	if err != nil {
		return core.CreateVcnResponse{}, errors.New("Unable to create vcn:" + err.Error())
	}

	return response, nil
}
