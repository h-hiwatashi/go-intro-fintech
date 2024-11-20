package main

import (
	// f"fmt"
	//do importは非推奨
	. "fmt"

	foo "github.com/h-hiwatashi/go-todolist/src/13/foo"
)

// scope

func TestScope(){
	//頭文字が大文字ならpablic
	Println(foo.Max)
	//頭文字が小文字の場合private
	// fmt.Println(foo.min)
	Println(foo.ReturnMin())
}

//関数内で定義された定数は他の関数で呼び出せない
// sは引数
// bは名前付きの戻り値
func TestScope2(s string)(b string){
	const a = "a"
	b = a + s
	{
		const c = "c"
		// 深いブロック内でのみ有効
		b := c + s
		Println(b)
	}
	Println(b)
	return
	
}

func main() {
	TestScope()
	TestScope2("AAA")
}