package cmd

import (
	"context"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/spf13/cobra"
)

var cidrBlock, name, dns string

var createNetworkVCNCmd = &cobra.Command{
	Use:   "vcn",
	Short: "create services(",
	Long:  `This subcommand allows you to create services`,

	Run: func(cmd *cobra.Command, args []string) {
		compartmentOCID := RootCmd.PersistentFlags().Lookup("compartment_ocid").Value.String()

		createVCN(compartmentOCID, cidrBlock, name, dns)
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
	tenant := TenancyOCID()
	user := UserOCID()
	region := Region()
	fingerprint := KeyFingerprint()
	privateKey := PrivateKey()
	config := common.NewRawConfigurationProvider(tenant, user, region, fingerprint, privateKey, nil)

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
