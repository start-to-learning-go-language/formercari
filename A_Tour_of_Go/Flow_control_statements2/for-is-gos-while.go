// セミコロンを省略することもできる。
// whileはGoではforだけを使う
package main

import "fmt"

func main()  {
  sum := 1
  for sum < 1000{
    sum += sum
  }
  fmt.Println(sum)
}
