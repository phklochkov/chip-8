package main

import (
  "fmt"
  "sortFactory"
  "panda"
)



func main() {
  mao := panda.New("Xiong")
  fmt.Println(mao)
  sortFactory.CreateSorter();
}
