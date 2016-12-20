package panda

import (
  "strconv"
)

// From fmt package, toString kind of stuff
type Stringer interface {
  String() string
}

type Panda struct {
  name string
  age int
}

func New(name string) Panda {
  return Panda{name: name}
}

// Implementation with correct receiver. Self-note: panda *Panda doesn't work, it should be passed by value
func (panda Panda) String() string {
  return panda.name + " it's " + strconv.Itoa(panda.age) + " years old"
}