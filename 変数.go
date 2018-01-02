var name string
name = "Mr. Go"
fmt.Println("Hello,", name)

var name = "Mr. Go"
fmt.Pringln("Hello,", name)

name := "Mr. Go"
fmt.Println("Hello,", name)

const title = "Go言語入門"
fmt.Println(title)

1234
093 //8進数の93
0xA3 //16進数のA3(「0XA3」と書いても良い)

3.1415
.25 //0.25
12. //12.0
1.25e-3 //0.00125の指数表記

2i
3.1415i
1.25e03i //指数表記も使用可能

'a'
'あ' //Unicodeであるためマルチバイト文字列もルーン1つで表現可能
'\n' //エスケープシーケンスを使用できる
'\u12AB' //コードポイントを直接記述可能
‘abc‘
‘\n‘ //改行ではなく\とnの二文字として扱われる。
‘ab
cd‘ //前の行と合わせて、改行を含む1つの文字列として扱われる。

"abc"
"abc\ncd"  //abとcdの間に改行が挿入される
"\u3042\u3044\u3046" //「あいう」のコードポイント表記