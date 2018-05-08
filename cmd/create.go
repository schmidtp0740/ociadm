package cmd

import (
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/spf13/cobra"
)

var createCMD = &cobra.Command{
	Use:   "create",
	Short: "create services(",
	Long:  `This subcommand allows you to create services`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		switch category := args[0]; category {
		case "network":
			switch something := args[1]; something {
			case "vcn":
				fmt.Println("create vcn")
			case "subnet":
				fmt.Println("create subnet")
			}
		case "compute":
			switch something := args[1]; something {
			case "instance":
				fmt.Println("create instance")
			case "image":
				fmt.Println("create image")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(createCMD)
}

func createVCN(config common.ConfigurationProvider, CidrBlock, CompartmentID, vcnDisplayName, dnsLabel string) core.Vcn {
	c, err := core.NewVirtualNetworkClientWithConfigurationProvider(config)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	request := core.CreateVcnRequest{}
	request.CidrBlock = common.String(CidrBlock)
	request.CompartmentId = common.String(CompartmentID)
	request.DisplayName = common.String(vcnDisplayName)
	request.DnsLabel = common.String(dnsLabel)

	r, err := c.CreateVcn(ctx, request)
	if err != nil {
		panic(err)
	}

	return r.Vcn
}
