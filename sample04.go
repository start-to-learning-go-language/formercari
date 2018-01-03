package main

import "fmt"

func main() {
  i ,sum := 1, 0
  for i <= 10000 {
    sum += i
    i ++
  }
  fmt.Println(sum)
}
