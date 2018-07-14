package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd ...
var RootCmd = &cobra.Command{
	Use:   "cobra-example",
	Short: "An example of cobra",
	Long:  `Manage OCI resources`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute ...
func Execute() {
	//compartmentID := os.Getenv("TF_VAR_compartment_ocid")

	if err := RootCmd.Execute(); err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

}

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// RootCmd.PersistentFlags().String("tenancy_ocid", viper.GetString("tenancy_ocid"), "tenant ocid")
	// RootCmd.PersistentFlags().String("user_ocid", viper.GetString("user_ocid"), "user ocid")
	// RootCmd.PersistentFlags().String("compartment_ocid", viper.GetString("compartment_ocid"), "tenant ocid")
	// RootCmd.PersistentFlags().String("fingerprint", viper.GetString("fingerprint"), "tenant ocid")
	// RootCmd.PersistentFlags().String("region", viper.GetString("region"), "tenant ocid")
	// RootCmd.PersistentFlags().String("private_key_path", viper.GetString("private_key_path"), "private_key_path")

	// RootCmd.MarkFlagRequired("tenancy_ocid")
	// RootCmd.MarkFlagRequired("user_ocid")
	// RootCmd.MarkFlagRequired("compartment_ocid")
	// RootCmd.MarkFlagRequired("fingerprint")
	// RootCmd.MarkFlagRequired("region")
	// RootCmd.MarkFlagRequired("private_key_path")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
