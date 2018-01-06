// 初期化と後処理ステートメントの記述は任意で良い。
package main

import "fmt"

func main()  {
  sum := 1
  for ; sum < 1000 ; {
    sum += sum
  }
  fmt.Println(sum)
}
