package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Printf("Hello World, from chip-8.\n")
  fmt.Println(Sqrt(3.), math.Sqrt(3.))
}

func Power(base float64, power int) float64 {
  result := 1.
  for i := 1; i <= power; i++ {
    result *= base
  }
  return result
}

func Sqrt(x float64) (float64) {
  start := x / 2.

  precision := 0.0001
  delta := start

  for delta > precision {
    var next float64 = start - (Power(start, 2) - x) / 2 * start

    delta = start - next

    if delta < 0 {
      delta = delta * (-1.)
    }

    start = next
  }

  return start
}
