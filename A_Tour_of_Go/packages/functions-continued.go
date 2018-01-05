//関数の2つ以上の引数が同じ型である場合には、最後の型を残して省略して記述できる。
//この例では
//x int, y int
//x, y int
package main

import "fmt"

func add(x, y int) int {
  return x + y
}

func main()  {
    fmt.Println(add(42, 13))
}
