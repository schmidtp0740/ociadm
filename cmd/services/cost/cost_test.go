package cost

import (
	"log"
	"testing"
)

// TestCost ...
func TestCost(t *testing.T) {
	err := Cmd.Execute()
	if err != nil {
		log.Fatal(err.Error())
	}
}
