package cmd

import (
	"context"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
)

func destroyVCN(config common.ConfigurationProvider, vcn core.Vcn) core.DeleteVcnResponse {
	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(config)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	request := core.DeleteVcnRequest{}
	request.VcnId = vcn.Id

	resp, err := client.DeleteVcn(ctx, request)
	if err != nil {
		panic(err)
	}
	return resp
}
