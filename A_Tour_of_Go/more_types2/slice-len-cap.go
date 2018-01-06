// スライスは長さと容量の療法を持っている
// スライスの長さは、それに含まれる要素の数です。
// スライスの容量は、スライスの最初の要素から数えて、元となる配列の要素です。
// スライスsの長さと容量はlen(s)とcap(s)という式を使用して得ることができます。
// 十分な容量を持って提供されているスライスを再スライスすることによって、スライスの長さを伸ばすことができる
package main

import "fmt"

func main()  {
  s := []int{2, 3, 5, 7, 11, 13}
  printSlice(s)

  //Slice the slice to give it zero length
  s = s[:0]
  printSlice(s)

  //Extend its length.
  s = s[:4]
  printSlice(s)

  // Drop its first two values.
  s = s[2:]
  printSlice(s)
}

func printSlice(s int[])  {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s),s)
}
