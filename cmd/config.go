package cmd

import (
	"io/ioutil"
)

// TenancyOCID ...
func TenancyOCID() string {
	return RootCmd.PersistentFlags().Lookup("tenancy_ocid").Value.String()
}

// UserOCID ...
func UserOCID() string {
	return RootCmd.PersistentFlags().Lookup("user_ocid").Value.String()
}

// KeyFingerprint ...
func KeyFingerprint() string {
	return RootCmd.PersistentFlags().Lookup("fingerprint").Value.String()
}

// Region ...
func Region() string {
	return RootCmd.PersistentFlags().Lookup("region").Value.String()
}

// PrivateKey ...
func PrivateKey() string {
	privateKeyPath := RootCmd.PersistentFlags().Lookup("private_key_path").Value.String()
	b, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		panic(err)
	}

	// block, _ := pem.Decode(privateKeyPath)
	// key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return string(b)
}
