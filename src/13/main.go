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

func main() {
	TestScope()
}