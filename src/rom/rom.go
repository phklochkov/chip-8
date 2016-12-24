package rom

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Rom struct {
	Data      []byte
	OpPointer uint16
}

func (rom *Rom) Size() int {
	return len(rom.Data)
}

func Load(path string) Rom {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	return Rom{Data: data}
}

func (obj *Rom) String() string {
	return string((*obj).Data)
}
