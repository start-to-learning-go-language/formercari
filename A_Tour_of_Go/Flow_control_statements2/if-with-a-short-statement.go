package main

import (
  "fmt"
  "math"
)
// ifステートメントは、forのように、条件の前に、評価するための簡単なステートメントをかける。
// ここで宣言された変数は、ifのスコープ内だけで有効
func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  }
  return lim
}

func main()  {
  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )
  return v
}
