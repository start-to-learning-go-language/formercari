package main

import "fmt"

func main() {
  sum := 1
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum)
}

//セミコロン;を省略することもできる
//while文はfor文によってGoでは使う
