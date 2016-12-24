package main

import (
	"chip"
	"rom"
)

func main() {
	game := rom.Load("D:/bo/Go/projects/chip-8/roms/BRIX")

	vm := chip.Chip{Rom: game}
	vm.Emulate()
}
