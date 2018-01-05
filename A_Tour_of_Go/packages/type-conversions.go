//型変換
// 変数v、型Tがあった場合、T(v)は、変数vをT型へ変換します。
// いくつかの変換を見てみましょう
// var i int = 42
// var f float64 = float64(i)
// var u uint = uint(f)
// よりシンプルに記述もできる
// i := 42
// f := float64(i)
// u := uint(f)

// Goでは型変換は明示的な変換が必要です。
package main

import (
  "fmt"
  "math"
)

func main()  {
  var x, y int = 3, 4
  var f float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(f)
  fmt.Println(x, y, z)
}
