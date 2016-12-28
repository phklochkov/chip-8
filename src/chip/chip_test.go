package chip_test

import (
	"chip"
	"rom"
	"testing"
)

func TestEmulate(t *testing.T) {
	// Should be 36 in total. In a perfect world
	testCases := [][]byte{
		[]byte{0x0A, 0xAA},
		[]byte{0x12, 0x34},
		[]byte{0x60, 0x42},
		[]byte{0x70, 0x42},
		[]byte{0x64, 0x1, 0x80, 0x40},
		[]byte{0x64, 0xF, 0x65, 0xF0, 0x84, 0x51},
		[]byte{0x64, 0xF, 0x65, 0xFF, 0x84, 0x52},
		[]byte{0x64, 0xF, 0x65, 0x0E, 0x84, 0x53},
		//[]byte{0x64, 0x42, 0x80, 0x43},
		//[]byte{0x64, 0x42, 0x80, 0x44},
		//[]byte{0x64, 0x42, 0x80, 0x45},
		//[]byte{0x64, 0x42, 0x80, 0x46},
		//[]byte{0x64, 0x42, 0x80, 0x47},
		//[]byte{0x64, 0x42, 0x80, 0x4E},
	}

	for _, data := range testCases {
		cartridge := rom.Rom{Data: data}
		vm := chip.Chip{Rom: cartridge}

		vm.Emulate()

		firstLeft := data[len(data)-2] >> 4
		lastLeft := data[len(data)-2] & 15

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
		case 6:
			if vm.VRegisters[lastLeft] != 0x42 {
				t.Error("6 is incorrect, register should be 42", data)
			}
		case 7:
			if vm.VRegisters[0] != 0x42 {
				t.Error("7XNN Fail")
			}
		case 8:
			firstRight := data[len(data)-1] >> 4
			lastRight := data[len(data)-1] & 0xF
			switch lastRight {
			case 0:
				if vm.VRegisters[lastLeft] != vm.VRegisters[firstRight] {
					t.Error("8XY0 Fail")
				}
			case 1:
				if vm.VRegisters[lastLeft] != 0xFF {
					t.Error("8XY1 Fail")
				}
			case 2:
				if vm.VRegisters[lastLeft] != 0xF {
					t.Error("8XY2 Fail")
				}
			case 3:
				if vm.VRegisters[lastLeft] != 1 {
					t.Error("8XY3 Fail")
				}
			case 4:
			case 5:
			case 6:
			case 7:
			case 0xE:
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
