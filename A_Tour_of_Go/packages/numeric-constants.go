// 数値の定数は、高精度な値(values)です。
//型のない定数は、その状況によって必要な型を取ることになります。
package main

import "fmt"

const (
  Big = 1 << 100
  Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1}
func needFloat(x float64) float64 {
  return x * 0.1
}

func main()  {
  fmt.Println(needInt(Small))
  fmt.Println(needFloat(Small))
  fmt.Println(needFloat(Big))
}
