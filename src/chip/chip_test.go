package chip_test

import (
	"chip"
	"rom"
	"testing"
)

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
	game := rom.Rom{Data: []byte{42, 42}}
	vm := chip.Chip{Rom: game}

	code, err := vm.NextOperation()

	if err != nil {
		t.Error(err)
	}
	if code[0] != code[1] && code[0] != 42 {
		t.Error("Code is retrieved incorrectly")
	}
	if vm.OpPointer != 2 {
		t.Error("Machine state is modified corrupt")
	}
}

func fix(t *testing.T) {
	if err := recover(); err != nil {
		t.Error("Something went horribly wrong:\n\t", err)
	}
}
