// 配列は固定長です。
// 一方で、スライスは可変長です。
// より柔軟な配列と見なすこともできます。
// 実際には、スライスは配列よりもより一般的です。
// 型[]Tは型Tのスライスを表します。
// はじめに5つの要素を持ったスライスを定義するには、次のように書く
// a[0:5]
package main

import "fmt"

func main()  {
  primes := [6]int{2, 3, 5, 7, 11, 13}

  var s []int = primes[1:4]
  fmt.Println(s)
}