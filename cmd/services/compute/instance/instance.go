package instance

import (
	"errors"
	"io/ioutil"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "instance",
	Short: "TODO",
	Long:  "TODO",
}

// GetComputeClientFromViper ...
func GetComputeClientFromViper() (core.ComputeClient, error) {
	tenancy := viper.GetString("default.tenancy")
	user := viper.GetString("default.user")
	region := viper.GetString("default.region")
	fingerprint := viper.GetString("default.fingerprint")
	privateKeyPath := viper.GetString("default.key_file")

	pemFileContent, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return core.ComputeClient{}, errors.New("Can not read private key location from config file")
	}
	privateKeyPassphrase := viper.GetString("default.privateKeyPassphrase")

	configProvider := common.NewRawConfigurationProvider(tenancy, user, region, fingerprint, string(pemFileContent), &privateKeyPassphrase)

	client, err := core.NewComputeClientWithConfigurationProvider(configProvider)
	if err != nil {
		return core.ComputeClient{}, errors.New("Virtual Compute Client Error:" + err.Error())
	}

	return client, nil
}
