package main

import (
  "fmt"
  "math"
)

func main() {
  var x, y int = 3, 4 // i := 42
  var f float64 = math.Sqrt(float64(x*x + y*y)) // f := float(i)
  var z unit = unit(f) // u := unit(f)
  fmt.Println(x, y, z)
}
