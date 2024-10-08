package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

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

//mathパッケージ
func testMath(){
	// math.Abs()
	// 絶対値を取得する
	fmt.Println(math.Abs(-1))
	// math.Ceil()
	// 小数点以下を切り上げる
	fmt.Println(math.Ceil(1.1))
	// math.Floor()
	// 小数点以下を切り捨てる
	fmt.Println(math.Floor(1.1))
	// math.Max()
	// 最大値を取得する
	fmt.Println(math.Max(1, 2))
	// math.Min()
	// 最小値を取得する
	fmt.Println(math.Min(1, 2))
	// math.Pow()
	// 累乗を計算する
	fmt.Println(math.Pow(2, 3))
	// math.Sqrt()
	// 平方根を計算する
	fmt.Println(math.Sqrt(4))
	//定数値一覧
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxUint8)
	fmt.Println(math.MaxUint16)
	fmt.Println(math.MaxUint32)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxInt)
	fmt.Println(math.Pi)

	//小数点切り捨て、切り上げ
	fmt.Println(math.Trunc(1.1))
	fmt.Println(math.Trunc(1.9))

	//FloorとCeilの違い
	fmt.Println(math.Floor(1.1))
	fmt.Println(math.Ceil(1.1))
	fmt.Println(math.Floor(1.9))
	fmt.Println(math.Ceil(1.9))
}


func main() {
	testMath()
}