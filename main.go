package main

import (
  "fmt"
  "strconv"
)

// From fmt package, ToString kind of stuff
type Stringer interface {
  String() string
}

type Panda struct {
  name string
  age int
}

// Implementation with correct receiver. Self-note: panda *Panda doesn't work, it should be passed by value
func (panda Panda) String() string {
  return panda.name + " it's " + strconv.Itoa(panda.age) + " years old"
}

func main() {
  xiong := Panda{
    name: "Xiong",
    age: 3,
  }

  fmt.Println("Hello World, from", xiong)
}
