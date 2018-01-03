package main

import(
  "fmt"
  "time"
)

func tak(x,y,z int) int {
  if x <= y {
    return z
  } else {
    return tak(tak(x - 1, y, z), tak(y - 1, z, x), tak(z - 1, x, y))
  }
}

func main() {
  s := time.Now()
  fmt.Println(tak(22, 11, 0))
  e := time.Now().Sub(s)
  fmt.Println(e)
}
