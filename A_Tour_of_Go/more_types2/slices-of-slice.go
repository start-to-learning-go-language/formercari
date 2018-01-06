//スライスは、他のスライスを含む任意の型を含むことができる。

package main

import (
  "fmt"
  "string"
)

func main()  {
  board := [][]string{
    []string{"_","_","_"},
    []string{"_","_","_"},
    []string{"_","_","_"},
  }

  board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

  for i := 0; i < len(board); i++ {
    fmt.Printf("%s\n", strings.Join(board[i]," "))
  }
}
