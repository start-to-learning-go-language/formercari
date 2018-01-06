// Goはポインタを使う
//var p *int
//&オペレーターは、そのオペランドへのポインタを引き出す。
//i := 42
//p = &i
//*オペレータは、ポインタの指す先の変数を示します。
//fmt.Println(*p)  ⬅︎ ポインタpを通じて、iから値を読み出す
//*p = 21 ⬅︎　ポインタpを通して、iへ値を代入する

package main

import "fmt"

func main()  {
  i, j := 42, 2701

  p := &i // iのアドレス
  fmt.Println(*p) //iの値
  *p = 21 //iの値に21を代入
  fmt.Println(i) //21

  p = &j //jのアドレス
  *p = *p / 37 //jを37でわる
  fmt.Println(j) //j
}
