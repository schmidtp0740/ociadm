package pkg

import (
	"errors"
	"io/ioutil"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/identity"
	"github.com/spf13/viper"
)

// GetComputeClient ...
// get customer config from viper for compute client
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

// GetDefaultComputeClient ...
// use default configuration provider for compute client
func GetDefaultComputeClient() (core.ComputeClient, error) {
	configProvider := common.DefaultConfigProvider()

	client, err := core.NewComputeClientWithConfigurationProvider(configProvider)
	if err != nil {
		return core.ComputeClient{}, errors.New("Virtual Compute Client Error:" + err.Error())
	}

	return client, nil
}

// GetNetworkClient ...
// use custom config file for network client
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

// GetDefaultNetworkClient ...
// use default configuration provider for compute client
func GetDefaultNetworkClient() (core.VirtualNetworkClient, error) {
	configProvider := common.DefaultConfigProvider()

	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(configProvider)
	if err != nil {
		return core.VirtualNetworkClient{}, errors.New("Virtual Network Client Error:" + err.Error())
	}

	return client, nil
}

// GetDefaultIdentityClient ...
func GetDefaultIdentityClient() (identity.IdentityClient, error) {
	configProvider := common.DefaultConfigProvider()

	client, err := identity.NewIdentityClientWithConfigurationProvider(configProvider)
	if err != nil {
		return identity.IdentityClient{}, errors.New("Identity  Client Error:" + err.Error())
	}

	return client, nil
}
