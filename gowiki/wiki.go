package main

import (
  "fmt"
  "io / ioutil"
)
//私たちは、fmt,ioutilパッケージをGo標準ライブラリからインポートする。
// 後で、追加の機能を実装するときに、このimport宣言にパッケージを追加する。

type Page struct {
  Title string
  Body []byte
}
