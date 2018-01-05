package main

import (
  "fmt"
  "math"
)

func sqrt(x float64) string {
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}

//Go言語のifセテートメントは、forループと同様に、()は不要で、{}は必要。

func main() {
  fmt.Println(sqrt(2), sqrt(-4))
}
