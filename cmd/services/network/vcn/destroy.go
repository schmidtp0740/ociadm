package vcn

import (
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/spf13/cobra"
)

// DestroyCmd destroy a vcn instance
var DestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "TODO",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {

		vcnID := args[0]

		destroyVCN(vcnID)
	},
}

func init() {
	Cmd.AddCommand(DestroyCmd)
}

func destroyVCN(vcnID string) core.DeleteVcnResponse {
	configProvider := common.ConfigurationProviderEnvironmentVariables("TF_VAR", "")

	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(configProvider)
	if err != nil {
		fmt.Println("error")
		panic(err)
	}

	ctx := context.Background()

	request := core.DeleteVcnRequest{}
	request.VcnId = common.String(vcnID)

	resp, err := client.DeleteVcn(ctx, request)
	if err != nil {
		fmt.Println("error")
		panic(err)
	}
	return resp
}
