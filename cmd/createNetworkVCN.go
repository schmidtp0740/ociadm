package cmd

import (
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/spf13/cobra"
)

var vcnCidrBlock, name, dns string

var createNetworkVCNCmd = &cobra.Command{
	Use:   "vcn",
	Short: "TODO",
	Long:  "TODO",

	Run: func(cmd *cobra.Command, args []string) {

		vcn := createVCN(ociDetails.compartmentID, vcnCidrBlock, name, dns)

		fmt.Println("created", vcn)
	},
}

func init() {
	createNetworkVCNCmd.Flags().StringVarP(&vcnCidrBlock, "cidr-block", "c", "10.0.0.0/16", "vcnCidrBlock")
	createNetworkVCNCmd.Flags().StringVarP(&name, "name", "n", "", "VCN Display Name")
	createNetworkVCNCmd.Flags().StringVarP(&dns, "dns", "d", "dnsLabel", "dnsLabel")
	createNetworkVCNCmd.MarkFlagRequired("name")
	createNetworkCmd.AddCommand(createNetworkVCNCmd)
}

func createVCN(compartmentOCID, vcnCidrBlock, name, dns string) core.Vcn {
	tenant := ociDetails.getTenancyOCID()
	user := ociDetails.getUserOCID()
	region := ociDetails.getRegion()
	fingerprint := ociDetails.getKeyFingerprint()
	privateKey := ociDetails.getPrivateKey()
	config := common.NewRawConfigurationProvider(tenant, user, region, fingerprint, privateKey, nil)

	c, err := core.NewVirtualNetworkClientWithConfigurationProvider(config)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	request := core.CreateVcnRequest{}

	request.CidrBlock = common.String(vcnCidrBlock)
	request.CompartmentId = common.String(compartmentOCID)
	request.DisplayName = common.String(name)
	request.DnsLabel = common.String(dns)

	r, err := c.CreateVcn(ctx, request)
	if err != nil {
		panic(err)
	}

	return r.Vcn
}
