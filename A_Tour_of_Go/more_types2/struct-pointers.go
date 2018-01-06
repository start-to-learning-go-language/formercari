// structのフィールドは、structのポインタを通してアクセスすることもできる。
// フィールドXを持つstructのポインタpがある場合、フィールドXにアクセスするには、(*p).Xのように書くことができる。
// しかし、この表記法は大変面倒ですので、Goでは代わりにp.Xと書くことができる。

package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main()  {
  v := Vertex{1, 2}
  p := &v
  p.X = 1e9
  fmt.Println(v)
}
