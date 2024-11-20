package main

import "fmt"

// 変数

func main() {
	var i int = 1;
	fmt.Println(i);

	var t, f bool = true, false;
	fmt.Println(t, f);

	var(
		i2 int = 2;
		s2 string = "hello";
	)
	fmt.Println(i2, s2);

	var i3 int;
	var s3 string;
	fmt.Println(i3, s3);

	// 暗黙的な定義は関数内のみ
	i4 := 4;
	fmt.Println(i4);

	// 明示的な定義は関数外でも可能
	var i5 int = 5;
	fmt.Println(i5);

	// 基本的には明示的な定義を使う
}