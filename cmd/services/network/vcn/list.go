package vcn

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/schmidtp0740/ociadm/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ListCmd ...
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "TODO",
	Long:  "TODO",

	Run: func(cmd *cobra.Command, args []string) {

		request, err := getListVcnRequestFromArgs(cmd, args)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		response, err := getListOfVcn(request)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		outputListVcnResponse(output, response)

	},
}

func init() {
	ListCmd.Flags().StringP("compartment-id", "c", "", "ocid of compartment")
	ListCmd.Flags().StringP("display-name", "n", "", "name of the vcn")
	ListCmd.Flags().Int("limit", 0, "number of results per page")
	ListCmd.Flags().StringP("lifecycle-state", "l", "", "lifecycle state")
	ListCmd.Flags().String("page", "", "token of 'opc-next-page'")

	// TODO add sort functionality
	// ListCmd.Flags().String("sort-order", "", "the sort order to use")
	// ListCmd.Flags().String("sort-by", "", "the field to sort by")
	Cmd.AddCommand(ListCmd)
}

func getListOfVcn(request core.ListVcnsRequest) (core.ListVcnsResponse, error) {

	tenancy := viper.GetString("default.tenancy")
	user := viper.GetString("default.user")
	region := viper.GetString("default.region")
	fingerprint := viper.GetString("default.fingerprint")
	privateKeyPath := viper.GetString("default.key_file")

	client, err := pkg.GetNetworkClient(tenancy, user, region, fingerprint, privateKeyPath)
	if err != nil {
		return core.ListVcnsResponse{}, err
	}

	ctx := context.Background()

	response, err := client.ListVcns(ctx, request)
	if err != nil {
		return core.ListVcnsResponse{}, errors.New("Unable to get list of vcns: " + err.Error())
	}

	return response, nil

}

func getListVcnRequestFromArgs(cmd *cobra.Command, args []string) (request core.ListVcnsRequest, err error) {

	request.CompartmentId = common.String(viper.GetString("default.compartment-id"))

	displayName, err := cmd.Flags().GetString("display-name")
	if err != nil {
		return core.ListVcnsRequest{}, err
	}

	if displayName != "" {
		request.DisplayName = common.String(displayName)
	}

	limit, err := cmd.Flags().GetInt("limit")
	if err != nil {
		return core.ListVcnsRequest{}, err
	}

	if limit > 0 {
		request.Limit = common.Int(limit)
	}

	page, err := cmd.Flags().GetString("page")
	if err != nil {
		return core.ListVcnsRequest{}, err
	}

	if page != "" {
		request.Page = common.String(page)
	}
	return request, nil
}

func outputListVcnResponse(output string, response core.ListVcnsResponse) {
	if output == "json" {
		var prettyJSON bytes.Buffer
		res := struct {
			Items       []core.Vcn `json:"data"`
			OpcNextPage *string    `json:"opcNextPage,omitempty"`
		}{
			Items:       response.Items,
			OpcNextPage: response.OpcNextPage,
		}
		responseAsBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = json.Indent(&prettyJSON, responseAsBytes, "", "  ")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("%s\n", string(prettyJSON.Bytes()))

	} else {

		fmt.Printf("%-30s%-20s%-80s\n", "Display Name", "CIDR Block", "OCID")
		for _, item := range response.Items {
			fmt.Printf("%-30s%-20s%-80s\n", *item.DisplayName, *item.CidrBlock, *item.Id)
		}

		if opcNextPage := response.OpcNextPage; opcNextPage != nil {
			fmt.Printf("\nopc-page-token: %s\n", *opcNextPage)
		}
	}
}
