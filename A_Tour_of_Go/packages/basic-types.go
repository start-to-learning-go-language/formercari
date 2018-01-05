//Go言語の基本型(組み込み型)は次の通りです。
// bool, string, int, int8 ,int64, uint, uint8,
package main

import (
  "fmt"
  "math/cmplx"
)

var (
  ToBe   bool           = false
  MaxInt uint64         = 1<<64 - 1
  z      complex128     = cmplx.Sqrt(-5 + 12i)
)

func main() {
  const f = "%T(%v)\n"
  fmt.Printf(f, ToBe, ToBe)
  fmt.Printf(f, MaxInt, MaxInt)
  fmt.Printf(f, z, z)
}
