package main

import "fmt"

// 型
// 整数型

func main() {
	// PCのbit数によってintのサイズが変わる
	// 32bitの場合は32bit、64bitの場合は64bit
	// 32bitの場合は-2147483648 ~ 2147483647
	// 64bitの場合は-9223372036854775808 ~ 9223372036854775807

	// 型を調べる方法
	var i int64 = 1;
	fmt.Println("%T\n", i);
}