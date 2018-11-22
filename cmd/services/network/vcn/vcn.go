package vcn

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
	Use:   "vcn",
	Short: "TODO",
	Long:  "TODO",
}

// GetNetworkClientFromViper ...
func GetNetworkClientFromViper() (core.VirtualNetworkClient, error) {
	tenancy := viper.GetString("default.tenancy")
	user := viper.GetString("default.user")
	region := viper.GetString("default.region")
	fingerprint := viper.GetString("default.fingerprint")
	privateKeyPath := viper.GetString("default.key_file")

	pemFileContent, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return core.VirtualNetworkClient{}, errors.New("Can not read private key location from config file")
	}
	privateKeyPassphrase := viper.GetString("default.privateKeyPassphrase")

	configProvider := common.NewRawConfigurationProvider(tenancy, user, region, fingerprint, string(pemFileContent), &privateKeyPassphrase)

	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(configProvider)
	if err != nil {
		return core.VirtualNetworkClient{}, errors.New("Vritual Network Client Error:" + err.Error())
	}

	return client, nil
}
