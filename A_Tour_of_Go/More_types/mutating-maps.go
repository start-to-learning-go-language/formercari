package main

import "fmt"

func main()  {
  m := make(map[string]int)

  m["Answer"] = 42
  fmt.Println("The value:", m["Answer"])

  m["Answer"] = 48
  fmt.Println("The value:", m["Answer"])

  delete(m, "Answer")
  fmt.Println("The value:", m["Answer"])

  v, ok := m["Answer"]
  fmt.Println("The value:", v, "Present?", ok)
}

//map m の操作を見ていきましょう。
//mへ要素(elem)の挿入や更新：
//m[key] = elem
//要素の取得:
//elem = m[key]
//要素の削除：
//delete(m, key)
