package pkg

import (
	"errors"
	"io/ioutil"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/spf13/viper"
)

// GetComputeClientFromViper ...
func GetComputeClient(tenancy, user, region, fingerprint, privateKeyPath string) (core.ComputeClient, error) {

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

// GetNetworkClientFromViper ...
func GetNetworkClient(tenancy, user, region, fingerprint, privateKeyPath string) (core.VirtualNetworkClient, error) {

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
