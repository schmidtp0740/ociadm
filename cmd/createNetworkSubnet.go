package cmd

import (
	"fmt"

	"github.com/oracle/oci-go-sdk/core"
	"github.com/spf13/cobra"
)

var availabilityDomain,
	subnetcidrBlock,
	compartmentID,
	dhcpOptionsID,
	displayName,
	dnsLabel,
	routeTableID string

var prohibitPublicIPOnVnic bool

var securityListIDs []string

var createNetworkSubnetCmd = &cobra.Command{
	Use:   "subnet",
	Short: "TODO",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		compartmentID = RootCmd.PersistentFlags().Lookup("compartment_ocid").Value.String()

		subnetResponse := createSubnet()

		fmt.Println("created", subnetResponse)
	},
}

func createSubnet() (subResponse core.CreateSubnetResponse) {

	return subResponse
}
