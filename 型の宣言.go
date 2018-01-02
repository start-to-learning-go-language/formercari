package main

import "fmt"

type Score int

func main() {
	var myScore Score = 100
	fmt.Printf("私の点数は%d点です。\n",myScore)
}

//可読性の向上

package main

import "fmt"

func main() {
	var readFunc func(struct{name string; meaning string}) string
	var dict struct{name string; meaning string}
	readFunc = readOut
	dict.name = "コーヒー"
	dict.meaning = "コーヒー豆から作られる黒色の飲み物"
	fmt.Println(readFunc(dict))
}

func readOut(s struct{name string; meaning string}) string {
	return fmt.Sprintf("「」は「」という意味です",s.name,s.meaning)
}

package main

import "fmt"

type Dictionary struct {
	name string
	meaning string
}

type ReadFunc func(Dictionary) string

func main() {
	var readFunc ReadFunc
	var dict Dictionary
	readFunc = readOut
	dict.name = "コーヒー"
	dict.meaning = "コーヒー豆から作られる黒色の飲み物"
	fmt.Println(readFunc(dict))
}

func readOut(d Dictionary) string {
	return fmt.Sprintf("「％s」は「%s」という意味です,d.name,d.meaning")
}

type Score int
func (s Score) Show() { fmt.Printf("点数は%dです\n",s)}
func main() {
	var myScore Score = 100
	myScore.Show()
}

package main

import "fmt"

type Score int

func main() {
	var myScore Score = 100
	showInt(int(myScore))
}

func showInt(i int) {
	fmt.Printf("value: %d\n", i)
}