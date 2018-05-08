package cmd

import (
	"context"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
)

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
