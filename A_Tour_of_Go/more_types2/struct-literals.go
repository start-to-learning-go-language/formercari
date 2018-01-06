// structリテラルは、フィールドの値を列挙することで新しいstructの初期値の割り当てを示す。
// Name:構文を使って、フィールドの一部だけを列挙することができる。（この方法でのフィールドの指定順序は関係ない)
// &を頭に付けると、新しく割り当てられたstructへのポインタを戻す。

package main

import "fmt"

type Vertex struct {
  X, Y int
}

var (
  v1 = Vertex{1, 2}
  v2 = Vertex{X: 1}
  v3 = Vertex{}
  p  = &Vertex{1, 2}
)

func main()  {
  fmt.Println(v1, p, v2, v3)
}
