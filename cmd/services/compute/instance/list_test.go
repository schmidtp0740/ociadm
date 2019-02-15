package instance

import (
	"fmt"
	"log"
	"testing"
)

func TestList(t *testing.T) {
	fmt.Println("testing oci compute instance list")
	if err := ListCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}

}
