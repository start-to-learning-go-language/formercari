//変数に初期値を与えずに宣言すると、ゼロ値(zero value)が与えられます。
// ゼロ値は型によって以下のように与えられる。
// 数値型(int,floatなど): 0
// bool型: false
// string型: ""(空文字列(empty string))

package main

import "fmt"

func main() {
  var i int
  var f float64
  var b bool
  var s string
  fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
