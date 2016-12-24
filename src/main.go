package main

import (
	"chip"
	"rom"
)

func main() {
	game := rom.Load("D:/bo/Go/projects/chip-8/roms/BRIX")

	vm := chip.Chip{}
	vm.Emulate(game)
}
