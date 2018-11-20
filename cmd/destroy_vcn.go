package cmd

import (
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/spf13/cobra"
)

func init() {
	var vcnId string

	var destroyVcnCmd = &cobra.Command{
		Use:   "vcn",
		Short: "TODO",
		Long:  "TODO",
		Run: func(cmd *cobra.Command, args []string) {

			destroyVCN(&vcnId)
		},
	}
	destroyVcnCmd.Flags().StringVarP(&vcnId, "vcn-id", "i", "", "OCID of vcn")
	destroyVcnCmd.MarkFlagRequired("vcn-id")
	DestroyCmd.AddCommand(destroyVcnCmd)
}

func destroyVCN(vcnId *string) core.DeleteVcnResponse {
	configProvider := common.ConfigurationProviderEnvironmentVariables("TF_VAR", "")

	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(configProvider)
	if err != nil {
		fmt.Println("error")
		panic(err)
	}

	ctx := context.Background()

	request := core.DeleteVcnRequest{}
	request.VcnId = vcnId

	resp, err := client.DeleteVcn(ctx, request)
	if err != nil {
		fmt.Println("error")
		panic(err)
	}
	return resp
}
