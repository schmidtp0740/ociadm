package cmd

import (
	"context"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/spf13/cobra"
)

var vcnID string

var destroyNetworkVCNCmd = &cobra.Command{
	Use:   "vcn",
	Short: "TODO",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		destroyVCN()
	},
}

func init() {
	destroyNetworkVCNCmd.Flags().StringVarP(&vcnID, "vcn-id", "i", "", "OCID of vcn")
	destroyNetworkVCNCmd.MarkFlagRequired("vcn-id")
	destroyNetworkCmd.AddCommand(destroyNetworkVCNCmd)
}

func destroyVCN() core.DeleteVcnResponse {
	tenant := TenancyOCID()
	user := UserOCID()
	region := Region()
	fingerprint := KeyFingerprint()
	privateKey := PrivateKey()
	config := common.NewRawConfigurationProvider(tenant, user, region, fingerprint, privateKey, nil)

	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(config)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	request := core.DeleteVcnRequest{}
	request.VcnId = common.String(vcnID)

	resp, err := client.DeleteVcn(ctx, request)
	if err != nil {
		panic(err)
	}
	return resp
}
