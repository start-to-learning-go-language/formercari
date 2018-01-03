//処理と条件式の間はセミコロン;で区切る。条件式を評価する前に処理が実行される。if文の中でしか使わない局所変数を定義するときに便利
if i := 0; i == 0 {
  fmt.Println("zero")
} else {
  fmt.Println("not zero")
}
