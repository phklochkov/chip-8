package chip

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
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
	for i := 0; i < 55; i++ {
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
			case 0xE0:
				cmd := exec.Command("cls")
				cmd.Stdout = os.Stdout
				cmd.Run()
				fmt.Println("Clear screen")
			case 0xEE:
				fmt.Println("Return")
			default:
				fmt.Println("Call RCA or something")
			}
		case 1:
			vm.Jump(opcode)
		case 2:
			fmt.Printf("Call subroutine - no idea")
		case 3:
			if vm.VRegisters[lastLeft] == opcode[1] {
				vm.Skip()
			}
		case 4:
			if vm.VRegisters[lastLeft] == opcode[1] {
				vm.Skip()
			}
		case 5:
			firstRight := opcode[1] >> 4
			if vm.VRegisters[lastLeft] == vm.VRegisters[firstRight] {
				vm.Skip()
			}
		case 6:
			vm.VRegisters[lastLeft] = opcode[1]
		case 7:
			vm.VRegisters[lastLeft] += opcode[1]
		case 8:
			firstRight := opcode[1] >> 4
			switch opcode[1] & 0xF {
			case 0:
				vm.VRegisters[lastLeft] = vm.VRegisters[firstRight]
			case 1:
				vm.VRegisters[lastLeft] |= vm.VRegisters[firstRight]
			case 2:
				vm.VRegisters[lastLeft] &= vm.VRegisters[firstRight]
			case 3:
				vm.VRegisters[lastLeft] ^= vm.VRegisters[firstRight]
			case 4:
				if firstRight > lastLeft {
					vm.VRegisters[0xF] = 1
				} else {
					vm.VRegisters[0xF] = 0
				}
				vm.VRegisters[lastLeft] += vm.VRegisters[firstRight]
			case 5:
				if firstRight > lastLeft {
					vm.VRegisters[0xF] = 0
				} else {
					vm.VRegisters[0xF] = 1
				}
				vm.VRegisters[lastLeft] -= vm.VRegisters[firstRight]
			case 6:
				vm.VRegisters[0xF] = vm.VRegisters[lastLeft] & 1
				vm.VRegisters[lastLeft] >>= 1
			case 7:
				if firstRight < lastLeft {
					vm.VRegisters[0xF] = 0
				} else {
					vm.VRegisters[0xF] = 1
				}
				vm.VRegisters[lastLeft] = vm.VRegisters[firstRight] - vm.VRegisters[lastLeft]
			case 0xE:
				vm.VRegisters[0xF] = vm.VRegisters[lastLeft] >> 7
				// WTF, Why on 0x8? Bollocks...
				vm.VRegisters[lastLeft] <<= 1
			}
		case 0x9:
			firstRight := opcode[1] >> 4
			if vm.VRegisters[lastLeft] != vm.VRegisters[firstRight] {
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
	if len((*vm).Rom.Data) == 0 {
		return nil, errors.New("You forgot to load ROM :)")
	}

	if len((*vm).Rom.Data)-int(vm.OpPointer+1) <= 0 {
		return nil, errors.New("The ROM has not enough data to fulfill the request")
	}

	opcode := []byte{(*vm).Rom.Data[vm.OpPointer], (*vm).Rom.Data[vm.OpPointer+1]}
	(*vm).OpPointer += 2

	return opcode, nil
}

func (vm *Chip) SetRegister(index, value byte) {
	(*vm).VRegisters[index] = value
}

func (vm *Chip) Skip() {
	vm.OpPointer += 2
}
