package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	bashCompletionFunc = `__oci_parse_get()
{
	local oci_output out
	if oci_output=$(oci-dev get --no-headers "$1" 2>/dev/null); then
		out=($(echo "${oci-dev_output}" | awk '{print $1}'))
		COMPREPLY=( $( compgen -W "${out[*]}" -- "$curl" ))
	fi
}

__oci_get_resource()
{
	if [[ ${#nouns[@]} -eq 0]]; then
		return 1
	fi
	__oci_parse__get ${nouns{#nouns[@]} -1}
	if [[ $? -eq 0 ]]' then
		return 0
	fi
}

__custom_func() {
	case ${last_command} in
		oci_get | oci_describe | oci_destroy )
		__oci_get_resource
		return
		;;
	*)
		;;
	esac
}
`
)

var cfgFile string

// RootCmd ...
var RootCmd = &cobra.Command{
	Use:   "oci",
	Short: "TODO",
	Long:  `TODO`,
	BashCompletionFunction: bashCompletionFunc,
}

// Execute ...
func Execute() {

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	// Executes every time command is run
	// cobra.OnInitialize(initConfig)

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().String("tenancy_ocid", viper.GetString("tenancy_ocid"), "tenant ocid")
	RootCmd.PersistentFlags().String("user_ocid", viper.GetString("user_ocid"), "user ocid")
	RootCmd.PersistentFlags().String("compartment_ocid", viper.GetString("compartment_ocid"), "tenant ocid")
	RootCmd.PersistentFlags().String("fingerprint", viper.GetString("fingerprint"), "tenant ocid")
	RootCmd.PersistentFlags().String("region", viper.GetString("region"), "tenant ocid")
	RootCmd.PersistentFlags().String("private_key_path", viper.GetString("private_key_path"), "private_key_path")

	RootCmd.MarkFlagRequired("tenancy_ocid")
	RootCmd.MarkFlagRequired("user_ocid")
	RootCmd.MarkFlagRequired("compartment_ocid")
	RootCmd.MarkFlagRequired("fingerprint")
	RootCmd.MarkFlagRequired("region")
	RootCmd.MarkFlagRequired("private_key_path")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}
