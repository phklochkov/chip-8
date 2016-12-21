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

		switch opcode {
		case 0x00E0:
			fmt.Println("Clear screen")
		default:
			fmt.Println("Coming soon")
		}
	}
}
