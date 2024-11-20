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

//構造体の埋め込み
type A struct {
	//フィールド名と型名が同じ場合は省略可能
	// User User
	User
}
func methodTest2(){
	a := A{User: User{Name: "tanaka", Age: 20}}
	fmt.Println(a)
	a.User.SayName()
	//クラスが省略されている時は、フィールド名でアクセスできる
	fmt.Println(a.Name)
	a.SayName()

	a.User.SetName2("suzuki")
	//エラーになる
	// User.SetName2("suzuki")
}

//構造体のコンストラクタ
func NewUser(name string, age int) *User {
	//アドレス演算子を使って、構造体のポインタを作成することが多い
	return &User{Name: name, Age: age}
}
func constructorTest(){
	user1 := NewUser("user1", 20)
	// ポインタ型であるため、アドレスが表示される
	fmt.Println(user1)
	// 実態にアクセスする場合は、アドレス演算子を使う
	fmt.Println(*user1)
}

// 構造体のスライス
//ポインタ型のスライスを定義するのが一般的
type Users []*User
func constructorSliceTest(){
	users1 := []*User{
		{Name: "user1", Age: 20},
		{Name: "user2", Age: 30},
		{Name: "user3", Age: 40},
	}
	fmt.Println(users1)
	user4 := User{Name: "user4", Age: 50}
	users1 = append(users1, &user4)
	fmt.Println(users1)
	for _, u := range users1 {
		fmt.Println(*u)
	}

	users2 := Users{}
	user5 := User{Name: "user5", Age: 60}
	users2 = append(users2, &user5)
	fmt.Println(users2)
}

// struct map
func structMapTest(){
	m1 := map[string]User{
		"user1": {Name: "user1", Age: 20},
		"user2": {Name: "user2", Age: 30},
		"user3": {Name: "user3", Age: 40},
	}
	fmt.Println(m1)

	m3 := make(map[string]User)
	fmt.Println(m3)
	m3["a"] = User{Name: "user1", Age: 20}
	fmt.Println(m3)
}

// struct Defined type（独自型）
type MyInt int
func uniqueTest(){
	var i1 MyInt
	fmt.Println(i1)
	fmt.Printf("%T\n", i1)

	i1.Double()
}
func (m MyInt) Double() MyInt {
	res := m +1
	fmt.Println(res)
	return res
}


func main() {
	uniqueTest()
}