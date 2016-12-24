package chip

import (
	"testing"
)

func TestGetMemoryLocation(t *testing.T) {
	twoByteOpcode := []byte{0xA3, 0x0C}
	if answer := GetMemoryLocation(twoByteOpcode); answer != 0x30C {
		t.Error("The mem-address is incorrect", answer)
	}

}
