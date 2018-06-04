package cmd

import (
	"io/ioutil"
)

type ociDetailStruct struct {
	tenancyOCID    string
	userOCID       string
	compartmentID  string
	fingerprint    string
	region         string
	privateKeyPath string
}

var ociDetails ociDetailStruct

// TenancyOCID ...
func (o *ociDetailStruct) getTenancyOCID() string {
	return o.tenancyOCID
}

// UserOCID ...
func (o *ociDetailStruct) getUserOCID() string {
	return o.userOCID
}

// KeyFingerprint ...
func (o *ociDetailStruct) getKeyFingerprint() string {
	return o.fingerprint
}

// Region ...
func (o *ociDetailStruct) getRegion() string {
	return o.region
}

func (o *ociDetailStruct) getPrivateKey() string {
	b, err := ioutil.ReadFile(o.privateKeyPath)
	if err != nil {
		panic(err)
	}

	return string(b)
}
