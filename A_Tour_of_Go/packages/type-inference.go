package main

import "fmt"

func main()  {
  v := 0.54i
  fmt.Printf("v is of type %T\n", v)
}

// 明示的な型を指定せずに変数を宣言する場合(:=やvar = のいずれか)
// 変数の型は右側の変数から型推論される。
// 右側の変数が型を持っている場合、左側の新しい変数は同じ型になる。
// var i int
// j := i // j is an int
