// [n]T型は、型Tのn個の変数の配列(array)を表します。
// 以下は,intの10個の配列を宣言しています。：
// var a [10]int
// 配列の長さは、型の一部分です。
// ですので、配列のサイズを変えることはできません。
// これは制約のように思えますが、心配しないでください。
// Goは配列を扱うための便利な方法を提供しています。

package main

import "fmt"

func main()  {
  var a [2]string
  a[0] = "Hello"
  a[1] = "World"
  fmt.Println(a[0], a[1])
  fmt.Println(a)

  primes := [6]int{2, 3, 5, 7, 11, 13}
  fmt.Println(primes)
}
