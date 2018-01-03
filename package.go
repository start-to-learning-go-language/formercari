package somepkg

var SomeVar int
var someVar2 int

func SomeFunc()  {
  SomeVar = 10
  someVar2 = 5
}

func someFunc2()  {
  SomeFunc()
}


package otherpkg

func OtherFunc() {
  SomeFunc() //エラー
  SomeVar = 5 //これもエラー
}

package otherpkg

import "somepkg"

func OtherFunc() {
  somepkg.SomeFunc()
  somepkg.SomeVar = 5
}

package otherpkg

import some "somepkg"

func OtherFunc() {
  some.SomeFunc()
  some.SomeVar = 5
}


package otherpkg

import . "somepkg"

func OtherFunc() {
  SomeFunc()
  SomeVar = 5
}

package otherpkg

import "somepkg"

func OtherFunc()  {
  somepkg.SomeFunc()
  somepkg.SomeVar = 5
}

package otherpkg

import "somepkg"

func OtherFunc()  {
  somepkg.SomeFunc2() //エラー
  somepkg.someVar2 = 5 //エラー
}
