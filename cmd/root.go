package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/schmidtp0740/goci/cmd/services/compute"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd ...
var RootCmd = &cobra.Command{
	Use:   "goci",
	Short: "TODO",
	Long:  "TODO",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {},
}

// Execute ...
func Execute() {

	if err := RootCmd.Execute(); err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

}

func init() {

	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.oci/config)")
	RootCmd.PersistentFlags().StringP("output", "o", "default", "specify how to output")

	// RootCmd.AddCommand(network.Cmd)
	RootCmd.AddCommand(compute.Cmd)

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra-example" (without extension).
		viper.AddConfigPath(home + "/.oci")
		viper.AddConfigPath(".")
		viper.SetConfigName("config")

	}

	viper.SetEnvPrefix("TF_VAR_")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Unable to read config file: ", viper.ConfigFileUsed())
	}
}
