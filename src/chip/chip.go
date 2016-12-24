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
}

func GetMemoryLocation(opcode []byte) uint16 {
	return (uint16(opcode[0])&15)<<8 | uint16(opcode[1])
}

func (proc *Chip) Emulate(game rom.Rom) {
	//var systemRegister uint16 = 0
	//fmt.Printf("\nNow running game\n\n%d\n\n", *game)
	for i := 0; i < 10; i++ {
		opcode, err := proc.NextOperation(game)
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
			proc.Jump(opcode)
		case 2:
			fmt.Printf("Call subroutine - no idea")
		case 3:
			fmt.Println("Skip next if eq")
		case 4:
			fmt.Println("Skip if neq")
		case 0xA:
			fmt.Printf("Set register to ")
			proc.SetSysRegister(GetMemoryLocation(opcode))
			fmt.Printf("%X\n", proc.OpPointer)
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
	fmt.Println("Jumping to mem-location ", machine.OpPointer)
}

func (vm *Chip) NextOperation(game rom.Rom) ([]byte, error) {
	(*vm).OpPointer += 2

	if int(vm.OpPointer) > len(game.Data)-3 {
		return nil, errors.New("The ROM has ended, STAHP")
	}

	opcode := []byte{game.Data[vm.OpPointer-2], game.Data[vm.OpPointer-1]}

	return opcode, nil
}
