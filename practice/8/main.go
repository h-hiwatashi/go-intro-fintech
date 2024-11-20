package main

import (
	"fmt"
	"time"
)


func main() {
	testGoroutine()
}

func testInt() {
	var x = interface{}(1);
	i, isInt := x.(int);
	fmt.Println(i, isInt);

	// 型アサーション
	// エラーが起きないようにするためには、以下のようにする
	f, isFloat64 := x.(float64);
	fmt.Println(f, isFloat64);


	switch v := x.(type) {
	case bool:
		fmt.Println("bool", v);
	case int:
		fmt.Println("int", v);
	case string:
		fmt.Println("string", v);
	default:
		fmt.Println("default", v);
	}
}

// ラベル付きfor
// ラベル付きforは、breakやcontinueで指定したラベルのfor文を抜けることができる
func testLabel() {
	LOOP:
	for {
		for {
			fmt.Println("loop");
			break LOOP;
		}
	}
}

// defer
// 関数の最後に実行される
// 複数のdeferがある場合は、後から宣言されたものから実行される
// 例えば、ファイルを開いたら必ず閉じるような処理に使う
func testDefer() {
	defer fmt.Println("foo");
	fmt.Println("bar");
}


// panicとrecoverはエラーハンドリングを実装することが推奨されているためあまり使用しない
// panic
// エラーを発生させる
// panicが発生すると、それ以降の処理は実行されない
// panicが発生すると、deferは実行される
// panicが発生すると、その関数は終了する
// panicが発生すると、その関数の呼び出し元にpanicが伝播する
// panicが発生すると、プログラム全体が終了する
func testPanic() {
	panic("panic");
}
// recover
// panicが発生したときに、panicを取り除く
// panicが発生してもプログラムが終了しないようにする
// panicが発生しても、その後の処理を実行する
func testRecover() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x);
		}
	}();
	panic("panic");
}

// 並行処理
// goをつけることで、関数を並行処理で実行することができる
// 並行処理を行うと、関数の実行が終わる前に次の処理に移るため、結果が変わる可能性がある
// 並行処理を行うと、goroutineが生成される
// goroutineは、スレッドよりも軽量で、数万個のgoroutineを生成することができる
// goroutineは、スレッドよりもメモリを消費しない
// goroutineは、スレッドよりも高速に生成することができる
// goroutineは、スレッドよりも高速にコンテキストスイッチを行うことができる
// goroutineは、スレッドよりも高速に終了することができる
func testGoroutine() {
	go fmt.Println("goroutine");
	fmt.Println("main");
	go sub()
	go sub()
	for {
		fmt.Println("main loop");
		time.Sleep(100 * time.Millisecond);
	}
	}
func sub() {
	for {
		fmt.Println("sub loop");
		time.Sleep(100 * time.Millisecond);
	}
}

// init
// main関数の前に実行される関数
func init() {
	fmt.Println("init")
}