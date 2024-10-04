package main

import (
	"fmt"
	"time"
)

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

func testSliceCopyTest(){
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 5)
	copy(s2, s1)
	fmt.Println(s2)
	
	// 参照渡し
	s3 := []int{1, 2, 3, 4, 5}
	s4 := s3
	s4[0] = 100
	fmt.Println(s3)
	
	// 値渡し
	s5 := []int{1, 2, 3, 4, 5}
	s6 := make([]int, 4)
	// nはコピーした要素数
	// 要素より多いと末尾が消される
	n := copy(s6, s5)
	s6[0] = 100
	fmt.Println(n, s5, s6)
}

// スライスのfor文
func testSliceFor(){
	s := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(s); i++{
		fmt.Println(s[i])
	}
	// rangeを使用する
	for i, v := range s{
		fmt.Println(i, v)
	}
	// valuesのみ取得
	for _, v := range s{
		fmt.Println(v)
	}
	// インデックスのみ取得
	for i:= range s{
		fmt.Println(i)
	}
}

//スライスの可変長引数
//要素の数を指定せずに引数を渡すことができる
func testSliceVariable(s ...int) int {
	n:=0
	for _,v:=range s{
		n+=v
	}
	return n
}

//マップ
//キーと値のペアを管理するためのデータ構造
func testMap(){
	// var m map[string]int
	// m = make(map[string]int)
	m := make(map[string]int)
	m["taguchi"] = 200
	m["fkoji"] = 300
	fmt.Println(m)
	fmt.Println(m["taguchi"])
	fmt.Println(m["fkoji"])
	fmt.Println(m["dotinstall"])
	v, ok := m["dotinstall"]
	fmt.Println(v, ok)
	v2, ok2 := m["taguchi"]
	fmt.Println(v2, ok2)
	delete(m, "taguchi")
	fmt.Println(m)
	
	mm := make(map[string]int)
	s, ok := mm["1"]
	//値が存在しない場合は、ゼロ値が返される
	//二つ目の戻り値で、値が存在するかどうかを確認できる(true or false)
	fmt.Println(s, ok)
	
	//mapの削除
	mmm := map[string]int{"日本": 100, "アメリカ": 200}
	delete(mmm, "日本")
	fmt.Println(mmm)
	
	//mapの長さ
	fmt.Println(len(mmm))
	
	//mapのfor文
	mmmm := map[string]int{"100": 100, "200": 200}
	for k, v := range mmmm{
		fmt.Println(k, v)
	}
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

		//チャネルのクローズ
		//チャネルは、クローズすることができる
		//クローズしたチャネルにデータを送信すると、パニックが発生する
		//クローズしたチャネルからデータを受信すると、データが取得できる
		//クローズしたチャネルからデータを受信すると、ゼロ値が取得できる
		func testChannelClose(){
			ch := make(chan int, 2)
			ch <- 1
			close(ch)
			//チャネルのバッファが空の場合、かつクローズの時okにfalseが返される
			i, ok := <- ch
			fmt.Println(i, ok)
			i2, ok := <- ch
			fmt.Println(i2, ok)

		// 	go recieverClose(ch)
		// 	for i := 0; i < 10; i++{
		// 		ch <- i
		// 	}
		// 	close(ch)
		// 	time.Sleep(2 * time.Second)
		}
		// func recieverClose(ch <- chan int){
		// 	for {
		// 		i, ok := <- ch
		// 		if ok == false{
		// 			break
		// 		}
		// 		fmt.Println(i)
		// 	}
		// }

		//チャネルのfor文
		//チャネルをfor文で処理することができる
		//チャネルがクローズされるまで、データの受信を繰り返す
		func testChannelFor(){
			ch1 := make(chan int, 3)
			ch1 <- 1
			ch1 <- 2
			ch1 <- 3
			// for文で使う時はcloseしないとエラーになる
			close(ch1)
			for i := range ch1{
				fmt.Println(i)
			}
		}


func main() {
	//可変長引数
	// fmt.Println(testSliceVariable(1,2,3,4,5))


	testChannelFor()
}