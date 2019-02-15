package instance

import (
	"fmt"
	"log"
	"testing"
)

func TestGet(t *testing.T) {
	fmt.Println("testing oci compute instance get")
	if err := GetCmd.E(); err != nil {
		log.Fatal(err.Error())
	}

}
