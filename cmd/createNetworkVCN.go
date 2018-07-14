package cmd

import (
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/spf13/cobra"
)

var cidrBlock, name, dns string

var createNetworkVCNCmd = &cobra.Command{
	Use:   "vcn",
	Short: "TODO",
	Long:  "TODO",

	Run: func(cmd *cobra.Command, args []string) {
		compartmentOCID := RootCmd.PersistentFlags().Lookup("compartment_ocid").Value.String()

		vcn := createVCN(compartmentOCID, cidrBlock, name, dns)

		fmt.Println("created", vcn)
	},
}

func init() {
	createNetworkVCNCmd.Flags().StringVarP(&cidrBlock, "cidr-block", "c", "10.0.0.0/16", "cidrBlock")
	createNetworkVCNCmd.Flags().StringVarP(&name, "name", "n", "", "VCN Display Name")
	createNetworkVCNCmd.Flags().StringVarP(&dns, "dns", "d", "dnsLabel", "dnsLabel")
	createNetworkVCNCmd.MarkFlagRequired("name")
	createNetworkCmd.AddCommand(createNetworkVCNCmd)
}

func createVCN(compartmentOCID, cidrBlock, name, dns string) core.Vcn {
	config := common.DefaultConfigProvider()

	c, err := core.NewVirtualNetworkClientWithConfigurationProvider(config)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	request := core.CreateVcnRequest{}

	request.CidrBlock = common.String(cidrBlock)
	request.CompartmentId = common.String(compartmentOCID)
	request.DisplayName = common.String(name)
	request.DnsLabel = common.String(dns)

	r, err := c.CreateVcn(ctx, request)
	if err != nil {
		panic(err)
	}

	return r.Vcn
}
