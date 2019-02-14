package cost

import (
	"fmt"
	"log"
	"testing"
)

// TestCost ...
func TestCost(t *testing.T) {
	fmt.Print("test")
	err := Cmd.Execute()
	if err != nil {
		log.Fatal(err.Error())
	}
}
