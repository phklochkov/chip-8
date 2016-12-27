package chip_test

import (
	"chip"
	"rom"
	"testing"
)

func TestEmulate(t *testing.T) {
	testCases := [][]byte{
		[]byte{0x0A, 0xAA},
		[]byte{0x12, 0x34},
	}

	for _, data := range testCases {
		cartridge := rom.Rom{Data: data}
		vm := chip.Chip{Rom: cartridge}

		vm.Emulate()

		firstLeft := data[0] >> 4
		//lastLeft := data[0] & 15

		switch firstLeft {
		case 0:
			if vm.OpPointer != 2 {
				t.Error("0 is unhandled")
			}
		case 1:
			if vm.OpPointer != 0x34 {
				t.Error("Jump is incorrect")
			}
		case 2:
		case 3:
			if vm.OpPointer != 4 {
				t.Error("Eq-Skip is incorrect")
			}
		default:
			t.Error("Unhandled test case")

		}
	}
}

func TestSetRegister(t *testing.T) {
	vm := chip.Chip{}
	if vm.SetRegister(0xF, 0xFF); vm.VRegisters[0xF] != 0xFF {
		t.Error("Register is not set", 0xF, 0xFF, vm.VRegisters)
	}
}

func TestSkip(t *testing.T) {
	vm := chip.Chip{}
	if vm.Skip(); vm.OpPointer != 2 {
		t.Error("Skip is incorrect")
	}
}

func TestSetSysRegister(t *testing.T) {
	vm := chip.Chip{}
	vm.SetSysRegister(0x200)

	if vm.SystemRegister != 0x200 {
		t.Error("System register is broken", vm.SystemRegister)
	}
}

func TestJump(t *testing.T) {
	vm := chip.Chip{}
	twoByteOpcode := []byte{0xA3, 0x0C}

	if vm.Jump(twoByteOpcode); vm.OpPointer != 0x30C-0x200 {
		t.Error("Jump corrupts vm state", vm.OpPointer)
	}
}

func TestGetMemoryLocation(t *testing.T) {
	twoByteOpcode := []byte{0xA3, 0x0C}
	if answer := chip.GetMemoryLocation(twoByteOpcode); answer != 0x30C {
		t.Error("The mem-address is incorrect", answer)
	}
}

func TestNextOperation(t *testing.T) {
	defer fix(t)

	testData := [][]byte{
		nil,
		[]byte{},
		[]byte{42},
		[]byte{42, 42},
		[]byte{42, 42, 42},
	}

	for _, data := range testData {
		game := rom.Rom{Data: data}
		vm := chip.Chip{Rom: game}

		code, err := vm.NextOperation()

		if data == nil && err == nil {
			t.Error("There's no ROM, Neo")
		}

		switch len(data) {
		case 0:
			if err == nil {
				t.Error("There's no ROM, Neo")
			}
		case 1:
			if err == nil {
				t.Error("There's no ROM, Neo")
			}
		case 2:
			if vm.OpPointer != 2 {
				t.Error("Wrong state modification")
			}
			if code[0] != 42 && code[1] != 42 {
				t.Error("Incorrect data retrieval")
			}
		case 3:
			if vm.OpPointer != 2 {
				t.Error("Wrong state modification")
			}
			if code[0] != 42 && code[1] != 42 {
				t.Error("Incorrect data retrieval")
			}
		default:
			t.Error("Unhandled test case")
		}
	}
	/*
		game := rom.Rom{Data: testData}
		vm := chip.Chip{Rom: game}

		code, err := vm.NextOperation()

		if err != nil {
			t.Error(err)
		}
		if code[0] != 42 {
			t.Error("NextOperation is incorrect")
		}
	*/
}

func fix(t *testing.T) {
	if err := recover(); err != nil {
		t.Error("Something went horribly wrong:\n\t", err)
	}
}
