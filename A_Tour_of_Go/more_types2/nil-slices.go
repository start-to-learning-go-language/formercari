// スライスのゼロ値はnilです。
// nilスライスは0の長さと容量を持っており、何の元となる配列も持っていません

package main

import "fmt"

func main()  {
  var s []int
  fmt.Println(s, len(s), cap(s))
  if s == nil {
    fmt.Println("nil!")
  }
}
