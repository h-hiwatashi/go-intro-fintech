package main

import (
	"fmt"
	"os"
	"time"
)


func main() {
	testTime()
}

//OSパッケージ
func testOSExit(){
	// os.Args
	// コマンドライン引数を保持する
	// os.Args[0]は、実行ファイルのパス
	// os.Args[1]は、第一引数
	os.Exit(1)
	fmt.Println("start")
}

func testOSFileOperation(){
	_, err := os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil{
		fmt.Println(err)
	}

	// Create()
	// Open()
	// 
}

//go build -o args main.go
func testOSArgs(){
	fmt.Println(os.Args)

	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
}

//timeパッケージ
func testTime(){
	// time.Now()
	// 現在時刻を取得する
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.Nanosecond())
	fmt.Println(t.Location())
	fmt.Println(t.Weekday())
	fmt.Println(t.YearDay())
	fmt.Println(t.Format("2006/01/02 15:04:05"))

	// 時刻を比較する
	t1 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	t2 := time.Date(2020, 1, 2, 0, 0, 0, 0, time.Local)
	fmt.Println(t1.Before(t2))

	// 指定時間のスリープ
	time.Sleep(1 * time.Second)
	fmt.Println("sleep")
}