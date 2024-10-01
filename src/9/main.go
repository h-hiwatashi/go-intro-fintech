package main

import (
	"fmt"
	"time"
)


func main() {
	testSlice()
}

// channel
//複数のゴルーチン間のデータの受け渡しを行うためのデータ構造
//チャネルは、データの送受信を行うための機能を提供する
//チャネルは、先入先出のキューとして機能する
//チャネルは、データの送受信を行うための演算子が用意されている
//チャネルは、要素数を持つことができる
//チャネルは、キャパシティを持つことができる
//チャネルは、キャパシティを越えるとデッドロックが発生する
func testChannel(){
	// ch1 = make(chan int)
	// ch2 := make(chan int)
	ch3 := make(chan int, 5)

	ch3 <- 1;
	ch3 <- 2;
	ch3 <- 3;
	fmt.Println(ch3)
	i := <- ch3
	fmt.Println(i)
	fmt.Println(<-ch3)
	ch3 <- 4;
	ch3 <- 5;
	ch3 <- 6;
	ch3 <- 7;
	ch3 <- 8; // キャパシティを越えるとデッドロックが発生する
	fmt.Println(ch3)
}

// チャネルの送受信
// チャネルは、データの送受信を行うための演算子が用意されている
func testChannelGoRoutine(){
	ch1 := make(chan int)
	ch2 := make(chan int)
	// go func(){
	// 	ch <- 1
	// }()
	// i := <- ch
	// fmt.Println(i)
	go reciever(ch1)
	go reciever(ch2)

	for i := 0; i < 10; i++{
		ch1 <- i
		ch2 <- i
		time.Sleep(10 * time.Millisecond)
	}
}
func reciever(ch <- chan int){
	for {
		i := <- ch
		fmt.Println(i)
	}
}

// スライスについて
// スライスは、配列の一部を参照するためのデータ構造
func testSlice(){
	// var s []int
	// s = make([]int, 5)
	s := make([]int, 5)
	fmt.Println(s)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	s[3] = 4
	s[4] = 5
	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[2])
	fmt.Println(s[3])
	fmt.Println(s[4])

	//暗黙的なスライスの宣言
	s2 := []int{1, 2, 3, 4, 5}
	fmt.Println(s2)

	//make()関数を使用してスライスを作成する
	s3 := make([]int, 5)
	fmt.Println(s3)

	//スライスの要素数を取得する
	fmt.Println(len(s3))

	//スライスのキャパシティを取得する
	fmt.Println(cap(s3))
}