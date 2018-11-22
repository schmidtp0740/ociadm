package vcn

import (
	"os"

	"github.com/spf13/cobra"
)

// GetCmd ...
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "TODO",
	Long:  "TODO",

	Run: func(cmd *cobra.Command, args []string) {

		// fmt.Println(viper.AllKeys())

		// listOfVcn := getVcn(compartmentId, displayName, lifeCycleState)

		// fmt.Printf("%-30s%-20s%-80s\n", "Display Name", "CIDR Block", "OCID")
		// for _, item := range listOfVcn {
		// 	fmt.Printf("%-30s%-20s%-80s\n", *item.DisplayName, *item.CidrBlock, *item.Id)
		// }

	},
}

func init() {
	var compartmentID, displayName, lifeCycleState string

	GetCmd.Flags().StringVarP(&displayName, "name", "n", "", "name of the vcn")
	GetCmd.Flags().StringVarP(&lifeCycleState, "lifecycle-state", "l", "", "lifecycle state")
	GetCmd.Flags().StringVarP(&compartmentID, "compartment-id", "i", os.Getenv("TF_VAR_compartment_ocid"), "ocid of compartment")
	Cmd.AddCommand(GetCmd)
}

// func getVcn(compartmentID, displayName, lifeCycleState string) []core.Vcn {

// 	ctx := context.Background()

// 	request := core.ListVcnsRequest{}

// 	request.CompartmentId = &compartmentID

// 	if displayName != "" {
// 		request.DisplayName = &displayName
// 	}

// 	if strings.ToLower(lifeCycleState) == "available" {
// 		request.LifecycleState = core.VcnLifecycleStateAvailable
// 	} else if strings.ToLower(lifeCycleState) == "provisioning" {
// 		request.LifecycleState = core.VcnLifecycleStateProvisioning
// 	} else if strings.ToLower(lifeCycleState) == "terminated" {
// 		request.LifecycleState = core.VcnLifecycleStateTerminated
// 	} else if strings.ToLower(lifeCycleState) == "terminating" {
// 		request.LifecycleState = core.VcnLifecycleStateTerminating
// 	}

// 	response, err := client.ListVcns(ctx, request)
// 	if err != nil {
// 		fmt.Println("error at create vcn")
// 		panic(err)
// 	}

// 	return response.Items
// }
