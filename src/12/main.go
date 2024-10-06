package main

import (
	"fmt"
)

// interface
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


func main() {
	interfaceTest()
}