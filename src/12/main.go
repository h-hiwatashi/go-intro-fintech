package main

import (
	"fmt"
)

// interface
// インターフェースは、メソッドの集まりを定義する型
// インターフェースは、具体的な型を隠蔽するために使われる
type Person struct {
	Name string
	Age int
}
func (p *Person) ToString() string{
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}
type Car struct {
	Number string
	Model string
}
func (c *Car) ToString() string{
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}
type Printable interface {
	//文字列を返すメソッド
	ToString() string
}

func interfaceTest(){
	vs := []Printable{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "XXX-0123", Model: "PX512"},
	}
	for _, v := range vs {
		fmt.Println(v.ToString())
	}
}

// interface カスタムエラー
type MyError struct {
	Message string
	ErrorCode int
}
func (e *MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.ErrorCode, e.Message)
}
func RaiseError() error {
	return &MyError{
		Message: "カスタムエラー",
		ErrorCode: 500,
	}
}

func errorTest() {
	err := &MyError{Message: "エラーが発生しました", ErrorCode: 404}
	fmt.Println(err.Error())

	raise := RaiseError()
	e, ok:= raise.(*MyError)
	if ok {
		fmt.Println(e.ErrorCode)
	}
}


// interface Stringer
type Stringer interface {
    String() string
}
// stringerについて



func main() {
	errorTest()
}