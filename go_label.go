package main

import "fmt"

func main()  {
  FOR_LABEL:
  for i := 0; i < 10; i++ {
    switch {
    case i == 3:
      // for文からの脱出
      break FOR_LABEL

    default:
      fmt.Println(i)
    }
  }
}

LABEL1:
for i := 0; i < 10; i++ {
  for j := 0; j < 10; j++ {
    //1番目のforへcontinue
    continue LABEL1
  } else if i == 1 && j == 1 {
    //2番目のforへcontinue
    continue
  }
}

For i := 0; i < 10; i++ {
  if i == 2 {
    //for文の外にあるLABELへ移動
    goto LABEL
  }
}
LABEL:
