package chip

import (
	"errors"
	"fmt"
	"rom"
)

type Chip struct {
	SystemRegister uint16
	VRegisters     [16]byte
	OpPointer      uint16
	Rom            rom.Rom
}

func GetMemoryLocation(opcode []byte) uint16 {
	return (uint16(opcode[0])&15)<<8 | uint16(opcode[1])
}
func (vm *Chip) Emulate() {
	//var systemRegister uint16 = 0
	//fmt.Printf("\nNow running game\n\n%d\n\n", *game)
	for i := 0; i < 10; i++ {
		opcode, err := vm.NextOperation()
		if err != nil {
			fmt.Println(err)
			break
		}

		firstLeft := opcode[0] >> 4
		lastLeft := opcode[0] & 15

		fmt.Printf("%X %X\n", (firstLeft<<4 | lastLeft), opcode[1])

		switch firstLeft {
		case 0:
			switch opcode[1] {
			case 0:
				fmt.Println("Clear screen")
			case 0xE:
				fmt.Println("Return")
			default:
				fmt.Println("Call RCA or something")
			}
		case 1:
			vm.Jump(opcode)
		case 2:
			fmt.Printf("Call subroutine - no idea")
		case 3:
			fmt.Println("Skip next if eq", lastLeft, opcode[1])
			if vm.VRegisters[lastLeft] == opcode[1] {
				vm.Skip()
			}
		case 4:
			fmt.Println("Skip if neq", lastLeft, opcode[1])
			if vm.VRegisters[lastLeft] == opcode[1] {
				vm.Skip()
			}
		case 5:
			firstRight := opcode[1] >> 4
			fmt.Println("Skip if VX == VY", lastLeft, firstRight)
			if vm.VRegisters[lastLeft] == vm.VRegisters[firstRight] {
				vm.Skip()
			}

		case 0xA:
			fmt.Printf("Set register to ")
			vm.SetSysRegister(GetMemoryLocation(opcode))
			fmt.Printf("%X\n", vm.OpPointer)
		default:
			fmt.Println("Coming soon")
		}
	}
}

func (vm *Chip) SetSysRegister(value uint16) {
	(*vm).SystemRegister = value
}

func (machine *Chip) Jump(opcode []byte) {
	// First 512 bytes (0x200 in hex) are chip-reserved memory
	machine.OpPointer = GetMemoryLocation(opcode) - 0x200
	fmt.Printf("Jumping to mem-location %X\n", machine.OpPointer)
}

func (vm *Chip) NextOperation() ([]byte, error) {
	(*vm).OpPointer += 2

	if int(vm.OpPointer) > len((*vm).Rom.Data) {
		return nil, errors.New("The ROM has ended, STAHP")
	}

	opcode := []byte{(*vm).Rom.Data[vm.OpPointer-2], (*vm).Rom.Data[vm.OpPointer-1]}

	return opcode, nil
}

func (vm *Chip) Skip() {
	vm.OpPointer += 2
}
