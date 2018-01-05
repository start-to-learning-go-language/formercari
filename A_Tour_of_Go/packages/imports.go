// このコードでは、括弧でパッケージのインポートをグループ化し、factoredインポートステートメントとしている。
// もちろん、複数のインポートステートメントで書くこともできる。
// import "fmt"
// import "math"

package main

import (
  "fmt"
  "math"
)

func main()  {
  fmt.Printf("Now you have %g problems.", math.Sqrt(7))
}
