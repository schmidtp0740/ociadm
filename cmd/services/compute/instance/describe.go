package instance

import (
	"github.com/spf13/cobra"
)

// DescribeCmd ...
var DescribeCmd = &cobra.Command{
	Use:   "describe [instance-id]",
	Short: "TODO",
	Long:  "TODO",

	Run: func(cmd *cobra.Command, args []string) {

		// instanceID := args[0]

		// instance := getInstanceDetails(instanceID)

		// fmt.Printf("Instance Display Name: %s\n", *instance.DisplayName)
		// fmt.Printf("Availability Domain %s\n", *instance.AvailabilityDomain)
		// fmt.Printf("Lifecycle State: %s\n", instance.LifecycleState)
		// fmt.Printf("Compartment OCID: %s\n", *instance.CompartmentId)
		// fmt.Printf("Instace OCID: %s\n", *instance.Id)
		// fmt.Printf("Region: %s\n", *instance.Region)
		// fmt.Printf("Shape: %s\n", *instance.Shape)
		// fmt.Printf("Time Created: %s\n", instance.TimeCreated.String())
		// fmt.Printf("Extended Metadata: %s\n", instance.ExtendedMetadata)
		// //fmt.Printf("Fault Domain: %s\n", *instance.FaultDomain)
		// fmt.Printf("ImageId: %s\n", *instance.ImageId)
		// fmt.Printf("Freeform tage: %s\n", instance.FreeformTags)
		// if instance.IpxeScript != nil {
		// 	fmt.Printf("Ipxe Script: %s\n", *instance.IpxeScript)

		// } else {
		// 	fmt.Printf("Ipxe Script: %s\n", "")

		// }
		// fmt.Printf("Launch Mode: %s\n", instance.LaunchMode)
		// fmt.Printf("Launch Options: %s\n", *instance.LaunchOptions)
		// fmt.Printf("Metadata: %s\n", instance.Metadata)
		// fmt.Printf("Source Details: %s\n", instance.SourceDetails)

		// // send ImageId get name of image
		// image := getImageName(*instance.ImageId)
		// fmt.Printf("Image Name: %s %s\n", *image.OperatingSystem, *image.OperatingSystemVersion)

		// // get list of vnics attached to instance
		// listOfVnicAttachments := listVnicAttachmentsForInstance(instanceID)

		// fmt.Printf("\nVNIC Attachment Details\n")
		// // for each vnic attachment
		// for _, vnicAttachment := range listOfVnicAttachments {
		// 	// print the details of the vnic attachment
		// 	// display name of vnic
		// 	if vnicAttachment.DisplayName != nil {
		// 		fmt.Printf("Display Name: %s\n", *vnicAttachment.DisplayName)
		// 	} else {
		// 		fmt.Printf("Display Name: %s\n", "")
		// 	}

		// 	// availability domain
		// 	fmt.Printf("Availability Domain: %s\n", *vnicAttachment.AvailabilityDomain)

		// 	// compartment OCID
		// 	fmt.Printf("Compartment OCID: %s\n", *vnicAttachment.CompartmentId)

		// 	// OCID of vnic attachement
		// 	fmt.Printf("VNIC Attachment OCID: %s\n", *vnicAttachment.Id)

		// 	// lifecycle state of vnic attachment
		// 	fmt.Printf("Life Cycle State: %s\n", vnicAttachment.LifecycleState)

		// 	// OCID of subnet
		// 	fmt.Printf("Subnet OCID: %s\n", *vnicAttachment.SubnetId)

		// 	// timecreated vnic attachment
		// 	fmt.Printf("Time Created: %s\n", *vnicAttachment.TimeCreated)

		// 	// nic index
		// 	fmt.Printf("NIC Index: %d\n", *vnicAttachment.NicIndex)

		// 	// vlan tag
		// 	fmt.Printf("VLAN Tag: %d\n", *vnicAttachment.VlanTag)

		// 	// get details of vnic
		// 	fmt.Printf("\nVNIC Details\n")

		// 	vnic := getVnicDetails(*vnicAttachment.VnicId)

		// 	// print details of vnic
		// 	// OCID of vnic
		// 	fmt.Printf("VLAN: %s\n", *vnic.Id)

		// 	// availability domain
		// 	fmt.Printf("Availability Domain: %s\n", *vnic.AvailabilityDomain)

		// 	// compartment OCID
		// 	fmt.Printf("Compartment OCID: %s\n", *vnic.CompartmentId)

		// 	// lifecycle state of vnic
		// 	fmt.Printf("Lifecycle State: %s\n", vnic.LifecycleState)

		// 	// private ip
		// 	fmt.Printf("Private IP Address: %s\n", *vnic.PrivateIp)

		// 	// OCID of subnet
		// 	fmt.Printf("Subnet OCID: %s\n", *vnic.SubnetId)

		// 	// time created vnic
		// 	fmt.Printf("Time Created: %s\n", *vnic.TimeCreated)

		// 	// defined tags
		// 	fmt.Printf("Defined Tags: %s\n", vnic.DefinedTags)

		// 	// display name of vnic
		// 	fmt.Printf("Display Name: %s\n", *vnic.DisplayName)

		// 	// freeform tags
		// 	fmt.Printf("Freeform Tags: %s\n", vnic.FreeformTags)

		// 	// Hostname label
		// 	fmt.Printf("Hosname Label: %s\n", *vnic.HostnameLabel)

		// 	// is the Vnic the primary vnic
		// 	fmt.Printf("Is Primary Vnic: %t\n", *vnic.IsPrimary)

		// 	// MacAddress
		// 	fmt.Printf("MAC Address: %s\n", *vnic.MacAddress)

		// 	// Public Ip
		// 	fmt.Printf("Public IP Address: %s\n", *vnic.PublicIp)

		// 	// Skip Source Destination check
		// 	fmt.Printf("Skip Source Destination Check: %t\n", *vnic.SkipSourceDestCheck)

		// }

		// get block volumes attached

		// for each block volume attachment
		//print details of block volume attachment
		//get details of block volume
		// print details of block volume

		// get boot volume details

		// get console connections details

	},
}

func init() {
	DescribeCmd.Flags().StringP("instanceOCID", "i", "", "OCID of instance")

	Cmd.AddCommand(DescribeCmd)
}

// func getDescribeInstanceRequestFromArgs(cmd *cobra.Command, args []string) (request core.GetInstanceRequest, err error) {

// 	instanceID := args[0]
// 	if err != nil {
// 		return core.GetInstanceRequest{}, err
// 	}

// 	request.InstanceId = common.String(instanceID)

// 	return request, nil
// }

// func getInstanceDetails(instanceID string) core.Instance {
// 	configProvider := common.ConfigurationProviderEnvironmentVariables("TF_VAR", "")

// 	client, err := core.NewComputeClientWithConfigurationProvider(configProvider)
// 	if err != nil {
// 		fmt.Print("Error: ")
// 		panic(err)
// 	}

// 	ctx := context.Background()

// 	request := core.GetInstanceRequest{}

// 	request.InstanceId = &instanceID

// 	response, err := client.GetInstance(ctx, request)
// 	if err != nil {
// 		fmt.Println("error at get instances")
// 		panic(err)
// 	}
// 	return response.Instance
// }

// func listVnicAttachmentsForInstance(instanceID string) []core.VnicAttachment {

// 	configProvider := common.ConfigurationProviderEnvironmentVariables("TF_VAR", "")

// 	client, err := core.NewComputeClientWithConfigurationProvider(configProvider)
// 	if err != nil {
// 		fmt.Print("Error: ")
// 		panic(err)
// 	}

// 	ctx := context.Background()

// 	request := core.ListVnicAttachmentsRequest{}

// 	// request.CompartmentId = &compOCID
// 	request.InstanceId = &instanceID

// 	response, err := client.ListVnicAttachments(ctx, request)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return response.Items
// }

// func getVnicDetails(VnicOCID string) core.Vnic {
// 	configProvider := common.ConfigurationProviderEnvironmentVariables("TF_VAR", "")

// 	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(configProvider)
// 	if err != nil {
// 		fmt.Print("Error: ")
// 		panic(err)
// 	}

// 	ctx := context.Background()

// 	request := core.GetVnicRequest{}

// 	request.VnicId = &VnicOCID

// 	response, err := client.GetVnic(ctx, request)
// 	if err != nil {
// 		fmt.Println("error at get instances")
// 		panic(err)
// 	}

// 	return response.Vnic

// }

// func listBlockVolumeAttachmentsForInstance(instanceID string) {

// }

// func getBlockVolumeDetails() {

// }

// func getImageName(imageID string) core.Image {
// 	configProvider := common.ConfigurationProviderEnvironmentVariables("TF_VAR", "")

// 	client, err := core.NewComputeClientWithConfigurationProvider(configProvider)
// 	if err != nil {
// 		fmt.Print("Error: ")
// 		panic(err)
// 	}

// 	ctx := context.Background()

// 	request := core.GetImageRequest{}

// 	request.ImageId = &imageID

// 	response, err := client.GetImage(ctx, request)
// 	if err != nil {
// 		fmt.Println("error at get instances")
// 		panic(err)
// 	}

// 	return response.Image

// }

// // func outputDescribeInstanceResponse(output string, response AllInstanceDetails) {
// // 	if output == "json" {
// // 		var prettyJSON bytes.Buffer
// // 		res := struct {
// // 			Data core.Instance `json:"data"`
// // 			Etag string        `json:"etag"`
// // 		}{
// // 			Data: response.Instance,
// // 			Etag: *response.Etag,
// // 		}
// // 		responseAsBytes, err := json.Marshal(res)
// // 		if err != nil {
// // 			fmt.Println(err.Error())
// // 			return
// // 		}
// // 		err = json.Indent(&prettyJSON, responseAsBytes, "", "  ")
// // 		if err != nil {
// // 			fmt.Println(err.Error())
// // 			return
// // 		}
// // 		fmt.Printf("%s\n", string(prettyJSON.Bytes()))

// // 	} else {

// // 	}
// // }
