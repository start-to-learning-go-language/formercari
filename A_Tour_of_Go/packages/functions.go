//関数は、0個以上の引数を取ることができます。
//変数名の後ろに型名を書くことに注意してください。
package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

func main()  {
  fmt.Println(add(42, 13))
}
