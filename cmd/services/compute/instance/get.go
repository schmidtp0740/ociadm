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
)

// GetCmd ...
var GetCmd = &cobra.Command{
	Use:   "get [instance-id]",
	Short: "TODO",
	Long:  "TODO",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		request, err := getInstanceRequestFromArgs(cmd, args)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		response, err := getInstance(request)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		outputGetInstanceResponse(output, response)

		// listOfInstances := getInstance(compartmentId, displayName, lifeCycleState, availabilityDomain)

		// fmt.Printf("%-50s%-20s%-20s%-80s\n", "Display Name", "Lifecycle State", "Shape", "OCID")
		// for _, item := range listOfInstances {
		// 	fmt.Printf("%-50s%-20s%-20s%-80s\n", *item.DisplayName, item.LifecycleState, *item.Shape, *item.Id)
		// }

	},
}

func init() {

	Cmd.AddCommand(GetCmd)
}

func getInstanceRequestFromArgs(cmd *cobra.Command, args []string) (request core.GetInstanceRequest, err error) {

	instanceID := args[0]
	if err != nil {
		return core.GetInstanceRequest{}, err
	}

	request.InstanceId = common.String(instanceID)

	return request, nil
}

func getInstance(request core.GetInstanceRequest) (core.GetInstanceResponse, error) {

	client, err := GetComputeClientFromViper()
	if err != nil {
		return core.GetInstanceResponse{}, err
	}

	ctx := context.Background()

	response, err := client.GetInstance(ctx, request)
	if err != nil {
		return core.GetInstanceResponse{}, errors.New("Unable to get instances: " + err.Error())
	}

	return response, nil

}

func outputGetInstanceResponse(output string, response core.GetInstanceResponse) {
	if output == "json" {
		var prettyJSON bytes.Buffer
		res := struct {
			Data core.Instance `json:"data"`
			Etag string        `json:"etag"`
		}{
			Data: response.Instance,
			Etag: *response.Etag,
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
		space := "30"
		fmt.Printf("%-"+space+"s%-80s\n", "Display Name:", *response.DisplayName)
		fmt.Printf("%-"+space+"s%-80s\n", "Availability Domain:", *response.AvailabilityDomain)
		fmt.Printf("%-"+space+"s%-80s\n", "Compartment ID:", *response.CompartmentId)
		fmt.Printf("%-"+space+"s%-80s\n", "Instance ID:", *response.Id)
		fmt.Printf("%-"+space+"s%-80s\n", "Lifecycle State:", response.LifecycleState)
		fmt.Printf("%-"+space+"s%-80s\n", "Region:", *response.Region)
		fmt.Printf("%-"+space+"s%-80s\n", "Shape:", *response.Shape)
		fmt.Printf("%-"+space+"s%-80s\n", "Time Created:", *response.TimeCreated)

		fmt.Printf("%-"+space+"s\n", "Extended Metadata:")
		for key, value := range response.ExtendedMetadata {
			fmt.Printf("  %-"+space+"s%-40s\n", key, value)
		}

		// Need to upgrade SDK for fault Domain
		// fmt.Printf("%-"+space+"s%-80s\n", "Fault Domain:", *response.FaultDomain)
		fmt.Printf("%-"+space+"s%-80s\n", "Image ID:", *response.ImageId)
		if response.IpxeScript != nil {
			fmt.Printf("%-"+space+"s%-80s\n", "Ipxe Script:", *response.IpxeScript)
		} else {
			fmt.Printf("%-"+space+"s%-80s\n", "Ipxe Script:", "nil")
		}
		fmt.Printf("%-"+space+"s%-80s\n", "Launch Mode:", response.LaunchMode)

		fmt.Printf("%-"+space+"s\n", "Launch Options:")
		fmt.Printf("  %-"+space+"s%-80s\n", "Boot Volume Type:", response.LaunchOptions.BootVolumeType)
		fmt.Printf("  %-"+space+"s%-80s\n", "Firmware:", response.LaunchOptions.Firmware)
		fmt.Printf("  %-"+space+"s%-80s\n", "Network Type:", response.LaunchOptions.NetworkType)
		fmt.Printf("  %-"+space+"s%-80s\n", "Remote Data Volume Type:", response.LaunchOptions.RemoteDataVolumeType)

		fmt.Printf("%-"+space+"s\n", "Metadata:")
		for key, value := range response.Metadata {
			fmt.Printf("  %-"+space+"s%-40s\n", key, value)
		}

		// Need to upgrade Source Details
		// fmt.Printf("%-"+space+"s\n", "Source Details:")

		// Need to upgrade SDK for maintenance reboot due
		// fmt.Printf("%-"+space+"s%-80s\n", "Time Maintenance Reboot Due:", *response.)
		fmt.Printf("%-"+space+"s%-80s\n", "Time Created:", *response.TimeCreated)

	}
}

// func getInstance(compartmentId, displayName, lifeCycleState, availabilityDomain string) []core.Instance {
// 	configProvider := common.ConfigurationProviderEnvironmentVariables("TF_VAR", "")

// 	client, err := core.NewComputeClientWithConfigurationProvider(configProvider)
// 	if err != nil {
// 		fmt.Print("Error: ")
// 		panic(err)
// 	}

// 	ctx := context.Background()

// 	request := core.ListInstancesRequest{}

// 	request.CompartmentId = &compartmentId

// 	if displayName != "" {
// 		request.DisplayName = &displayName
// 	}

// 	if availabilityDomain != "" {
// 		request.AvailabilityDomain = &availabilityDomain
// 	}

// 	if strings.ToLower(lifeCycleState) == "running" {
// 		request.LifecycleState = core.InstanceLifecycleStateRunning
// 	} else if strings.ToLower(lifeCycleState) == "provisioning" {
// 		request.LifecycleState = core.InstanceLifecycleStateProvisioning
// 	} else if strings.ToLower(lifeCycleState) == "stopped" {
// 		request.LifecycleState = core.InstanceLifecycleStateStopped
// 	} else if strings.ToLower(lifeCycleState) == "terminated" {
// 		request.LifecycleState = core.InstanceLifecycleStateTerminated
// 	} else if strings.ToLower(lifeCycleState) == "starting" {
// 		request.LifecycleState = core.InstanceLifecycleStateStarting
// 	} else if strings.ToLower(lifeCycleState) == "stopping" {
// 		request.LifecycleState = core.InstanceLifecycleStateStopping
// 	} else if strings.ToLower(lifeCycleState) == "terminating" {
// 		request.LifecycleState = core.InstanceLifecycleStateTerminating
// 	}

// 	response, err := client.ListInstances(ctx, request)
// 	if err != nil {
// 		fmt.Println("error at get instances")
// 		panic(err)
// 	}

// 	return response.Items
// }
