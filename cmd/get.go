package cmd

import (
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:       "get",
	Short:     "TODO",
	ValidArgs: []string{"network"},
	Long:      "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "network" {
			getNetworks()
		}
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}

func getNetworks() {
	config := common.NewRawConfigurationProvider(ociDetails.tenancyOCID,
		ociDetails.userOCID,
		ociDetails.region,
		ociDetails.fingerprint,
		ociDetails.getPrivateKey(),
		nil,
	)

	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(config)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	request := core.ListVcnsRequest{
		CompartmentId: common.String(ociDetails.compartmentID),
	}

	response, err := client.ListVcns(ctx, request)
	if err != nil {
		panic(err)
	}

	fmt.Println("get networks")
	fmt.Println(response)

}
