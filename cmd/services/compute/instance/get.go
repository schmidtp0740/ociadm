package instance

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/schmidtp0740/goci/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type fullInstanceDetails struct {
	Instance        core.Instance
	Image           core.Image
	VnicAttachments []core.VnicAttachment
	Vnic            []core.Vnic
	BootVolume      core.BootVolume
	Etag            *string `json:"etag,omitempty"`
}

// GetCmd ...
var GetCmd = &cobra.Command{
	Use:   "get [instance-id]",
	Short: "TODO",
	Long:  "TODO",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		instance, err := getInstance(cmd, args[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		outputGetInstanceResponse(cmd, instance)

	},
}

func init() {

	GetCmd.Flags().BoolP("all", "a", false, "get all attachment details to instance")

	Cmd.AddCommand(GetCmd)
}

func getInstance(cmd *cobra.Command, instanceID string) (fullInstanceDetails, error) {

	fullInstance := fullInstanceDetails{}

	allFlag, err := cmd.Flags().GetBool("all")
	if err != nil {
		return fullInstanceDetails{}, err
	}

	tenancy := viper.GetString("default.tenancy")
	user := viper.GetString("default.user")
	region := viper.GetString("default.region")
	fingerprint := viper.GetString("default.fingerprint")
	privateKeyPath := viper.GetString("default.key_file")

	client, err := pkg.GetComputeClient(tenancy, user, region, fingerprint, privateKeyPath)
	if err != nil {
		return fullInstanceDetails{}, err
	}

	ctx := context.Background()

	getInstanceRequest := core.GetInstanceRequest{}

	getInstanceRequest.InstanceId = common.String(instanceID)

	getInstanceResponse, err := client.GetInstance(ctx, getInstanceRequest)
	if err != nil {
		return fullInstanceDetails{}, errors.New("Unable to get instances: " + err.Error())
	}

	fullInstance.Instance = getInstanceResponse.Instance

	fullInstance.Etag = getInstanceResponse.Etag

	if allFlag == true {

		// Get Image Details

		getInstanceImageRequest := core.GetImageRequest{}

		getInstanceImageRequest.ImageId = fullInstance.Instance.ImageId

		getInstanceImageResponse, err := client.GetImage(ctx, getInstanceImageRequest)
		if err != nil {
			fmt.Println("error at get instances")
			panic(err)
		}

		fullInstance.Image = getInstanceImageResponse.Image

		// get Vnic Attachment Details

		listVnicAttachmentRequest := core.ListVnicAttachmentsRequest{}

		listVnicAttachmentRequest.InstanceId = fullInstance.Instance.Id

		listVnicAttachmentResponse, err := client.ListVnicAttachments(ctx, listVnicAttachmentRequest)
		if err != nil {
			log.Fatal(err)
		}

		fullInstance.VnicAttachments = listVnicAttachmentResponse.Items

		// Get VNIC details for each VNIC attachment to the instance

		for _, vnicAttachment := range fullInstance.VnicAttachments {

			networkClient, err := pkg.GetNetworkClient(tenancy, user, region, fingerprint, privateKeyPath)
			if err != nil {
				return fullInstanceDetails{}, err
			}
			getVnicRequest := core.GetVnicRequest{}

			getVnicRequest.VnicId = vnicAttachment.VnicId

			getVnicResponse, err := networkClient.GetVnic(ctx, getVnicRequest)
			if err != nil {
				fmt.Println("error at get instances")
				panic(err)
			}

			fullInstance.Vnic = append(fullInstance.Vnic, getVnicResponse.Vnic)
		}

	}

	return fullInstance, nil

}

func outputGetInstanceResponse(cmd *cobra.Command, instance fullInstanceDetails) {

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if output == "json" {
		var prettyJSON bytes.Buffer

		res := struct {
			Data fullInstanceDetails `json:"data"`
		}{}

		res.Data = instance

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
		fmt.Printf("%-"+space+"s%-80s\n", "Display Name:", instance.Instance.DisplayName)
		fmt.Printf("%-"+space+"s%-80s\n", "Availability Domain:", instance.Instance.AvailabilityDomain)
		fmt.Printf("%-"+space+"s%-80s\n", "Compartment ID:", instance.Instance.CompartmentId)
		fmt.Printf("%-"+space+"s%-80s\n", "Instance ID:", instance.Instance.Id)
		fmt.Printf("%-"+space+"s%-80s\n", "Lifecycle State:", instance.Instance.LifecycleState)
		fmt.Printf("%-"+space+"s%-80s\n", "Region:", instance.Instance.Region)
		fmt.Printf("%-"+space+"s%-80s\n", "Shape:", instance.Instance.Shape)
		fmt.Printf("%-"+space+"s%-80s\n", "Time Created:", instance.Instance.TimeCreated)

		fmt.Printf("%-"+space+"s\n", "Extended Metadata:")
		for key, value := range instance.Instance.ExtendedMetadata {
			fmt.Printf("  %-"+space+"s%-40s\n", key, value)
		}

		// Need to upgrade SDK for fault Domain
		// fmt.Printf("%-"+space+"s%-80s\n", "Fault Domain:", *response.FaultDomain)
		fmt.Printf("%-"+space+"s%-80s\n", "Image ID:", instance.Instance.ImageId)
		if instance.Instance.IpxeScript != nil {
			fmt.Printf("%-"+space+"s%-80s\n", "Ipxe Script:", instance.Instance.IpxeScript)
		} else {
			fmt.Printf("%-"+space+"s%-80s\n", "Ipxe Script:", "nil")
		}
		fmt.Printf("%-"+space+"s%-80s\n", "Launch Mode:", instance.Instance.LaunchMode)

		fmt.Printf("%-"+space+"s\n", "Launch Options:")
		fmt.Printf("  %-"+space+"s%-80s\n", "Boot Volume Type:", instance.Instance.LaunchOptions.BootVolumeType)
		fmt.Printf("  %-"+space+"s%-80s\n", "Firmware:", instance.Instance.LaunchOptions.Firmware)
		fmt.Printf("  %-"+space+"s%-80s\n", "Network Type:", instance.Instance.LaunchOptions.NetworkType)
		fmt.Printf("  %-"+space+"s%-80s\n", "Remote Data Volume Type:", instance.Instance.LaunchOptions.RemoteDataVolumeType)

		fmt.Printf("%-"+space+"s\n", "Metadata:")
		for key, value := range instance.Instance.Metadata {
			fmt.Printf("  %-"+space+"s%-40s\n", key, value)
		}

		// Need to upgrade Source Details
		// fmt.Printf("%-"+space+"s\n", "Source Details:")

		// Need to upgrade SDK for maintenance reboot due
		// fmt.Printf("%-"+space+"s%-80s\n", "Time Maintenance Reboot Due:", *response.)
		fmt.Printf("%-"+space+"s%-80s\n", "Time Created:", instance.Instance.TimeCreated)

	}
}
