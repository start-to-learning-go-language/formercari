//関数の中では、var宣言の代わりに、短い:=の代入文を使い、暗黙的な型宣言ができます。
//なお、関数の外では、キーワードで始まる宣言(var,func,など)が必要で、:=での暗黙的な宣言は利用できません。
package main

import "fmt"

func main()  {
  var i, j int = 1, 2
  k := 3
  c, python, java := true, false, "no!"

  fmt.Println(i, j, k, c, python, java)
}
