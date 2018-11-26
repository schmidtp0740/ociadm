package instance

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ListCmd ...
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "TODO",
	Long:  "TODO",

	Run: func(cmd *cobra.Command, args []string) {

		request, err := getListInstancesRequestFromArgs(cmd, args)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		response, err := getListOfInstances(request)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		outputListInstancesResponse(output, response)

	},
}

func init() {
	ListCmd.Flags().StringP("compartment-id", "c", "", "ocid of compartment")
	ListCmd.Flags().StringP("availability-domain", "a", "", "availability domain")

	ListCmd.Flags().StringP("display-name", "n", "", "name of the vcn")
	ListCmd.Flags().Int("limit", 0, "number of results per page")
	ListCmd.Flags().StringP("lifecycle-state", "l", "", "lifecycle state")
	ListCmd.Flags().String("page", "", "token of 'opc-next-page'")

	// TODO add sort functionality
	// ListCmd.Flags().String("sort-order", "", "the sort order to use")
	// ListCmd.Flags().String("sort-by", "", "the field to sort by")
	Cmd.AddCommand(ListCmd)
}

func getListInstancesRequestFromArgs(cmd *cobra.Command, args []string) (request core.ListInstancesRequest, err error) {

	request.CompartmentId = common.String(viper.GetString("default.compartment-id"))

	displayName, err := cmd.Flags().GetString("display-name")
	if err != nil {
		return core.ListInstancesRequest{}, err
	}

	if displayName != "" {
		request.DisplayName = common.String(displayName)
	}

	availabilityDomain, err := cmd.Flags().GetString("availability-domain")
	if err != nil {
		return core.ListInstancesRequest{}, err
	}

	if availabilityDomain != "" {
		request.AvailabilityDomain = common.String(availabilityDomain)
	}

	limit, err := cmd.Flags().GetInt("limit")
	if err != nil {
		return core.ListInstancesRequest{}, err
	}

	if limit > 0 {
		request.Limit = common.Int(limit)
	}

	page, err := cmd.Flags().GetString("page")
	if err != nil {
		return core.ListInstancesRequest{}, err
	}

	if page != "" {
		request.Page = common.String(page)
	}
	return request, nil
}

func getListOfInstances(request core.ListInstancesRequest) (core.ListInstancesResponse, error) {

	client, err := GetComputeClientFromViper()
	if err != nil {
		return core.ListInstancesResponse{}, err
	}

	ctx := context.Background()

	response, err := client.ListInstances(ctx, request)
	if err != nil {
		return core.ListInstancesResponse{}, errors.New("Unable to get list of vcns: " + err.Error())
	}

	return response, nil

}

func outputListInstancesResponse(output string, response core.ListInstancesResponse) {
	if output == "json" {
		var prettyJSON bytes.Buffer
		res := struct {
			Items       []core.Instance `json:"data"`
			OpcNextPage *string         `json:"opcNextPage,omitempty"`
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

		fmt.Printf("%-42s%-80s\n", "Display Name", "OCID")
		for _, item := range response.Items {
			fmt.Printf("%-42s%-80s\n", *item.DisplayName, *item.Id)
		}

		if opcNextPage := response.OpcNextPage; opcNextPage != nil {
			fmt.Printf("\nopc-page-token: %s\n", *opcNextPage)
		}
	}
}
