// Goのプログラムは、パッケージで構成される。
//プログラムはmainパッケージから開始される。
//このプログラムでは"fmt"と"math/rand"パッケージをインポート(import)している。
//規約で、パッケージ名はインポートパスの最後の要素と同じ名前になる。
//例えば、インポートパスが"math/rand"のパッケージは、package randステートメントで始まるファイル群で構成する。
package main

import (
  "fmt"
  "math/rand"
)

func main()  {
  fmt.Println("My favorite number is", rand.Intn(10))
}
