//解説
//開始点zを選び、式のループを繰り返すことによって、Sqrt(x)を近似していく。
//１、最初は10回だけ繰り返し、xを(1,2,3,...)と様々な値に対する結果がどれだけ正解値に近づくかを確認する。
//

package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {//xの値を(1,2,3,...)と色々試してみる
  z := float64(2.)
  s := float64(0)
  for i := 0; i < 10; i++ {
    z = z - (z*z - x) / (2*z)
    if math.Abs(z - s) < 1e-10 {
      break;//２、ループを回す時の直前に求めたzの値がこれ以上変化しなくなった時(または、差がとても小さくなった時)に停止するようにループを変更してみてください。
      //この変更により、ループ回数が多くなったか、少なくなったのかをみてみる。(math.Sqrt)と比べてどれくらい近似できたのか調べる。
    }
    s = z
  }
  return z
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(math.Sqrt(2))
}
