package main

import "fmt"

func main()  {
  fmt.Println("counting")

  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }
  fmt.Println("done")
}

// deferへ渡した関数が複数ある場合、その呼び出しはスタック(stack)される。
// 呼び出し元の関数がreturnするとき、deferへ渡した関数はlast-in-first-outの順番で実行される。
