package main

import "fmt"

type User struct {
	//定義がない場合はゼロ値が入る
	Name string
	Age int
	// x, Y int
}
func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 3000
}
func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 3000
}
func structTest(){
	var user1 User
	fmt.Println(user1)

	user1.Name = "tanaka"
	user1.Age = 20
	fmt.Println(user1)

	user2 := User{Name: "suzuki", Age: 30}
	fmt.Println(user2)

	user4 := User{"sato", 40}
	fmt.Println(user4)

	//7と8は同じ
	// アドレス演算子を使って、構造体のポインタを作成することが多い
	user7 := new(User)
	fmt.Println(user7)

	user8 := &User{}
	fmt.Println(user8)

	UpdateUser(user1)
	UpdateUser2(user8)
	// 値渡しのため、変更されない
	fmt.Println(user1)
	// 参照渡しのため、変更される
	fmt.Println(user8)
}

//構造体のメソッド
func (u User) SayName() {
	fmt.Println(u.Name)
}
func (u User) SetName(name string) {
	u.Name = name
}
func (u *User) SetName2(name string) {
	u.Name = name
}
func methodTest() {
	user1 := User{Name: "tanaka", Age: 20}
	user1.SayName()

	user1.SetName("suzuki")
	user1.SayName()

	//基本はポインタで渡す
	user1.SetName2("sato")
	user1.SayName()

	//ポインタ型で渡してもOK
	user2 := &User{Name: "tanaka", Age: 20}
	user2.SetName2("suzuki")
	user2.SayName()
}



func main() {
	methodTest()
}