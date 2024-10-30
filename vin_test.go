package vin

import (
	"testing"
)

const testVIN = "W09000051T2123456"

func TestVIN_Manufacturer(t *testing.T) {

	manufacturer := Manufacturer(testVIN)
	if manufacturer != "W09123" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
	}
}
