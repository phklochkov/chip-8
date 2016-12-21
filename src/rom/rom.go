package rom

import (
	"fmt"
	"os"
	"io/ioutil"
	"errors"
)

type Rom struct {
	Data []byte
	OpPointer int
}

func (rom *Rom) Size() int {
	return len(rom.Data)
}

func (rom *Rom) NextOperation() (byte, error){
	rom.OpPointer++
	if rom.OpPointer > len((*rom).Data) - 2 {
		return byte(0), errors.New("The ROM has ended, STAHP")
	}
	return rom.Data[rom.OpPointer - 1], nil
}

func Load(path string) *Rom {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	return &Rom{Data: data}
}

func (obj *Rom) String() string {
	return string(obj.Data)
}
