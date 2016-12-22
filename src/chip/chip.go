package chip

import (
	"fmt"
	"rom"
)

func Emulate(game *rom.Rom) {
	//fmt.Printf("\nNow running game\n\n%d", *game)
	for {
		opcode, err := game.NextOperation()
		if err != nil {
			break
		}
		
		first4 := opcode[0] >> 4
		second4 := opcode[0] & 15

		fmt.Printf("%b %b %b ", first4, second4, opcode)

		switch first4 {
		case 0:
			fmt.Println("Clear screen")
		default:
			fmt.Println("Coming soon")
		}
	}
}
