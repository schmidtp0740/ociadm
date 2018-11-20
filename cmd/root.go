package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd ...
var RootCmd = &cobra.Command{
	Use:   "oci-dev",
	Short: "TODO",
	Long:  "TODO",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute ...
func Execute() {

	if err := RootCmd.Execute(); err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

}

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// RootCmd.PersistentFlags().String("tenancy_ocid", os.Getenv("TF_VAR_tenancy_ocid"), "tenant ocid")
	// RootCmd.PersistentFlags().String("user_ocid", os.Getenv("TF_VAR_user_ocid"), "user ocid")
	// RootCmd.PersistentFlags().String("compartment_ocid", os.Getenv("TF_VAR_compartment_ocid"), "tenant ocid")
	// RootCmd.PersistentFlags().String("fingerprint", os.Getenv("TF_VAR_fingerprint"), "tenant ocid")
	// RootCmd.PersistentFlags().String("region", os.Getenv("TF_VAR_region"), "tenant ocid")
	// RootCmd.PersistentFlags().String("private_key_path", os.Getenv("TF_VAR_private_key_path"), "private_key_path")

	// RootCmd.MarkFlagRequired("tenancy_ocid")
	// RootCmd.MarkFlagRequired("user_ocid")
	// RootCmd.MarkFlagRequired("compartment_ocid")
	// RootCmd.MarkFlagRequired("fingerprint")
	// RootCmd.MarkFlagRequired("region")
	// RootCmd.MarkFlagRequired("private_key_path")
}
