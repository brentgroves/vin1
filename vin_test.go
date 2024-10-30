package vin

import (
	"testing"
)

const (
	validVIN   = VIN("W0L000051T2123456")
	invalidVIN = VIN("W0")
)

func TestVIN_Manufacturer(t *testing.T) {

	manufacturer := validVIN.Manufacturer()
	if manufacturer != "W0L" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, validVIN)
	}

	invalidVIN.Manufacturer() // panic!
}
