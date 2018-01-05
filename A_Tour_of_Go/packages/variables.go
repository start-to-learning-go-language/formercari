//varステートメントは変数を宣言する。
//関数の引数リストと同様に、複数の変数の最後に型を書くことで、変数のリストを宣言できる。
//varステートメントはパッケージ、または、関数で利用できる。
package main

import "fmt"

var c, python, java bool

func main()  {
  var i int
  fmt.Println(i, c, python, java)
}
