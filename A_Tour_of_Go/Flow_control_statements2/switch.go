//Goのswitchは、caseの最後で自動的にbreakする。
// もしbreakせずに通したい場合は、fallthrough文をcaseの最後に記述する。

package main

import (
  "fmt"
  "math"
)

func main()  {
  fmt.Print("Go runs on")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("Linux.")
  default:
    fmt.Printf("%s.", os)
  }
}
