package main

import "fmt"
//インデックスや値は、"_"(アンダーバー)へ代入することで捨てることができる。
func main() {
  pow := make([]int, 10)
  for i := range pow {
    pow[i] = 1 << uint(i)
  }
  for _, value := range pow {
    fmt.Printf("%d\n", value)
  }
}
