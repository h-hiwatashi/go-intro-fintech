package main

import (
	"fmt"
	"time"
)


func main() {
	testChannelGoRoutine()
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