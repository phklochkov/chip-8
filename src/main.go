package main

import (
	"rom"
	"chip"
)

func main() {
	game := rom.Load("D:/bo/Go/projects/chip-8/roms/BRIX")

	chip.Emulate(game)
}
