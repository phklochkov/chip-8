package chip

import (
	"fmt"
	"rom"
)

func GetMemoryLocation(opcode []byte) uint16 {
	return (uint16(opcode[0])&15)<<8 | uint16(opcode[1])
}

func Emulate(game *rom.Rom) {
	var systemRegister uint16 = 0
	//fmt.Printf("\nNow running game\n\n%d\n\n", *game)
	for i := 0; i < 10; i++ {
		opcode, err := game.NextOperation()
		if err != nil {
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
			case 14:
				fmt.Println("Return")
			default:
				fmt.Println("Call RCA or something")
			}
		case 1:
			// First 512 bytes (0x200 in hex) are chip-reserved memory
			game.OpPointer = GetMemoryLocation(opcode) - 0x200
			fmt.Println("Jumping to mem-location ", game.OpPointer)
		case 2:
			fmt.Printf("Call sub")
		case 3:
			fmt.Println("Skip next if eq")
		case 4:
			fmt.Println("Skip if neq")
		case 5:
			fmt.Println("Clear screen")
		case 6:
			fmt.Println("Clear screen")
		case 7:
			fmt.Println("Clear screen")
		case 8:
			fmt.Println("Clear screen")
		case 9:
			fmt.Println("Clear screen")
		case 0xA:
			fmt.Printf("Set register to ")
			systemRegister = GetMemoryLocation(opcode)
			fmt.Printf("%X\n", systemRegister)
		case 11:
			fmt.Println("Clear screen")
		case 12:
			fmt.Println("Clear screen")
		case 13:
			fmt.Println("Clear screen")
		case 14:
			fmt.Println("Clear screen")
		case 15:
			fmt.Println("Clear screen")

		default:
			fmt.Println("Coming soon")
		}
	}
}
