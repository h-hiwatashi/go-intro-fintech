package main

import (
	"fmt"
)

//ポインタについて
//ポインタは、変数のアドレスを格納するためのデータ構造
func testPointer(){
	// 値渡し


	var i int = 100
	doubeleTest(i)
	// ここのiは100のまま
	fmt.Println(i)

	// &iはiのアドレス演算子
	// pはiのアドレスを格納するための変数
	// 参照渡し
	var p *int = &i
	*p = 200
	fmt.Println(p)
	// *pはpのアドレスに格納されている値
	//　200になる
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(i)
	

	var sls []int = []int{1, 2, 3, 4, 5}
	// 参照渡しで全て2倍にする
	doubeleTest3(sls)
	fmt.Println(sls)
}
func doubeleTest(i int) {
	i = i * 2
}
func doubeleTest3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	testPointer()
}