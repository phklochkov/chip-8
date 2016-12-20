package sortFactory

import "fmt"

type Sortable interface {
  Len() int
}

func CreateSorter() {
  fmt.Println("Created")
}