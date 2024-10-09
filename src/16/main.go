package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
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

// randomパッケージ
func testRandom(){
	// 0からn未満の乱数を生成する
	// rand.Intn()
	// 乱数のシードを設定する
	// rand.Seed()
	// 現在の時間をシードに設定する
	// rand.Seed(time.Now().UnixNano())
	// 0からn未満の乱数を生成する
	// fmt.Println(rand.Intn(100))

	// 推奨の書き方
	// 現在の時間をシードに設定して新しい乱数生成器を作成する
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 0からn未満の乱数を生成する
	fmt.Println(r.Intn(100))
}

// flagパッケージ
func testFlag(){
//コマンドラインの引数をパースする
// go run main.go -n -m message -x

	// flagパッケージ
	// コマンドライン引数をパースする
	// flag.BoolVar()
	// flag.IntVar()
	// flag.StringVar()
	// flag.Parse()
	// flag.Args()
	// flag.NArg()
	// flag.Arg(i)
	// flag.Usage()
	// flag.Visit()
	// flag.VisitAll()
	var (
		i int
		b bool
		s string
	)
	flag.IntVar(&i, "n", 32, "処理の最大値")
	flag.BoolVar(&b, "b", false, "真偽値")
	flag.StringVar(&s, "s", "aaa", "文字列")

	// パースする
	flag.Parse()

	fmt.Println("処理の最大値=", i)
	fmt.Println("真偽値=", b)
	fmt.Println("文字列=", s)
}

// fmtパッケージ
func testFmt(){
	fmt.Println("改行")
	fmt.Printf("フォーマット\n")

	// 書き込み先を指定する
	// Fprint: フォーマットなしで文字列を出力します。
	// Fprintf: フォーマットされた文字列を出力します。
	// Fprintln: フォーマットなしで文字列を出力し、最後に改行を追加します。
	// fmt.Fprint()
	// fmt.Fprintf()
	// fmt.Fprintln()

	//　出力ではなく文字列を返す
	// fmt.Sprint()
	// fmt.Sprintf()
	// fmt.Sprintln()

	
	// fmt.Errorf()

	// 指定されたio.Readerから入力を読み取り
	// fmt.Fscan()
	// fmt.Fscanf()
	// fmt.Fscanln()

	// fmt.Scanf()
	// fmt.Sscan()
	// fmt.Sscanln()
	}

func main() {
	testFlag()
}