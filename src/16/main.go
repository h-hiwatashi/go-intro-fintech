package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
	// fmt.Println("改行")
	// fmt.Printf("フォーマット\n")

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


	// verb
	// 書式指定子
	// %v: デフォルトフォーマット
	// 様々な型の値を柔軟に出力する
	fmt.Printf("%v\n", true)
	fmt.Printf("%v\n", 1)
	fmt.Printf("%v\n", 1.1)
	fmt.Printf("%v\n", "1")
	fmt.Printf("%v\n", []int{1, 2, 3})

	// %+v: 構造体のフィールド名を表示
	// %#v: Goの構文形式で表示

	// %T: 型を表示
	fmt.Printf("%T\n", true)
	fmt.Printf("%T\n", 1)
	fmt.Printf("%T\n", 1.1)
	fmt.Printf("%T\n", "1")
	fmt.Printf("%T\n", []int{1, 2, 3})
	// %t: 真偽値を表示

	// %b: 2進数を表示
	// %c: Unicodeコードポイントを表示
	// %d: 10進数を表示
	// %o: 8進数を表示
	// %x: 16進数を表示（小文字）
	// %X: 16進数を表示（大文字）
	// %e: 科学技術表記で表示（小文字）
	// %E: 科学技術表記で表示（大文字）
	// %f: 浮動小数点数を表示
	// %g: %eまたは%fの短い形式を表示
	// %G: %Eまたは%fの短い形式を表示
	// %s: 文字列を表示
	// %q: クォートされた文字列を表示
	// %p: ポインタを表示
	// %U: Unicodeフォーマットを表示
	}


	//logパッケージ
	// log.Fatal()
	func testLog(){
		log.SetOutput(os.Stdout)

		// ログメッセージを出力する
		// log.Print("ログメッセージ\n")
		// log.Println("ログメッセージ")
		// log.Printf("log message %d\n", 1)
		
		// logメッセージを出力し処理が終了する
		// log.Fatal("ログメッセージ\n")
		// log.Fatalln("ログメッセージ")
		// log.Fatalf("log message %d\n", 1)

		// panicメッセージを出力し処理が終了する
		// パニックは、エラーが発生した場合にプログラムを終了する
		// log.Panic("ログメッセージ\n")
		// log.Panicln("ログメッセージ")
		// log.Panicf("log message %d\n", 1)

		// ログの出力先を指定する
		// 任意のファイルにログを出力することができる
		// ファイルが存在しない場合は、新規作成される
		// f, err := os.Create("test.log")
		// if err != nil {
		// 	return
		// }
		// log.SetOutput(f)
		// log.Println("ログメッセージ")


		// エラーメッセージを出力してプログラムを終了する
		// 標準ログのフォーマット
		log.SetFlags(log.LstdFlags)
		log.Println("A")
		// 2024/10/13 00:01:13 A

		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Println("B")
		// 2024/10/13 00:01:13.946518 B

		log.SetFlags(log.Ltime | log.Lshortfile)
		log.Println("C")
		// 00:01:13 main.go:281: C

		log.SetFlags(log.LstdFlags)
		log.SetPrefix("[INFO]")
		log.Println("D")
		//[INFO]2024/10/13 00:01:13 D



		// ロガーを作成する
		// ロガーを作成することで、出力先やフォーマットを指定することができる
		// 第一引数に出力先を指定する
		// 第二引数にプレフィックスを指定する
		// 第三引数にフラグを指定する
		logger := log.New(os.Stdout, "[ERROR]", log.LstdFlags)
		logger.Println("E")
		log.Println("F")
		// 条件分岐によってエラーの内容を出力する
		_, err := os.Open("test.txt")
		if err != nil {
			logger.Println(err)
		}
	}


	// strconvパッケージ
func testStrconv(){
	// strconv.Itoa()
	// 数値を文字列に変換する
	fmt.Println(strconv.Itoa(100))
	// strconv.Atoi()
	// 文字列を数値に変換する
	i, _ := strconv.Atoi("100")
	fmt.Println(i)
	// strconv.ParseInt()
	// 文字列を整数に変換する
	i64, _ := strconv.ParseInt("100", 10, 64)
	fmt.Println(i64)
	fmt.Println(i)
	// strconv.FormatInt()
	// 整数を文字列に変換する
	fmt.Println(strconv.FormatInt(100, 10))
	// strconv.ParseFloat()
	// 文字列を浮動小数点数に変換する
	f, _ := strconv.ParseFloat("1.1", 64)
	fmt.Println(f)
	// strconv.FormatFloat()
	// 浮動小数点数を文字列に変換する
	fmt.Println(strconv.FormatFloat(1.1, 'f', -1, 64))
	// strconv.ParseBool()
	// 文字列を真偽値に変換する
	b, _ := strconv.ParseBool("true")
	fmt.Println(b)
	// strconv.FormatBool()
	// 真偽値を文字列に変換する
	fmt.Println(strconv.FormatBool(true))
	// strconv.Quote()
	// 文字列をクォートする
	fmt.Println(strconv.Quote("test"))
	// strconv.QuoteToASCII()
	// ASCII文字列をクォートする
	fmt.Println(strconv.QuoteToASCII("test"))
	// strconv.QuoteRune()
	// ルーンをクォートする
	fmt.Println(strconv.QuoteRune('a'))
	// strconv.QuoteRuneToASCII()
	// ASCIIルーンをクォートする
	fmt.Println(strconv.QuoteRuneToASCII('a'))
	// strconv.Unquote()
	// クォートされた文字列をアンクォートする
	fmt.Println(strconv.Unquote("\"test\""))
	// strconv.AppendInt()
	// 整数を文字列に追加する
	byteSlice := []byte("test")
	byteSlice = strconv.AppendInt(byteSlice, 100, 10)
}

// Stringsパッケージ
func testStrings(){
	// 文字列の結合
	// strings.Join()
	// 文字列を結合する


	// strings.Contains()
	// 文字列が含まれているか判定する
	fmt.Println(strings.Contains("test", "es"))
	// strings.Count()
	// 文字列が含まれている数を取得する
	fmt.Println(strings.Count("test", "t"))
	// strings.HasPrefix()
	// 文字列が指定した文字列で始まっているか判定する
	fmt.Println(strings.HasPrefix("test", "te"))
	// strings.HasSuffix()
	// 文字列が指定した文字列で終わっているか判定する
	fmt.Println(strings.HasSuffix("test", "st"))
	// strings.Index()
	// 文字列が含まれている位置を取得する
	fmt.Println(strings.Index("test", "s"))
	// strings.LastIndex()
	// 文字列が含まれている最後の位置を取得する
	fmt.Println(strings.LastIndex("test", "t"))
	// strings.Replace()
	// 文字列を置換する
	fmt.Println(strings.Replace("test", "t", "T", -1))
	// strings.Split()
	// 文字列を分割する
	fmt.Println(strings.Split("test", "e"))
	// strings.ToLower()
	// 文字列を小文字に変換する
	fmt.Println(strings.ToLower("TEST"))
	// strings.ToUpper()
	// 文字列を大文字に変換する
	fmt.Println(strings.ToUpper("test"))
	// strings.TrimSpace()
	// 文字列の前後の空白を削除する
	fmt.Println(strings.TrimSpace(" test "))
	// strings.Trim()
	// 文字列の前後の指定した文字列を削除する
	fmt.Println(strings.Trim("test", "t"))
	// strings.TrimLeft()
	// 文字列の前の指定した文字列を削除する
	fmt.Println(strings.TrimLeft("test", "t"))
	// strings.TrimRight()
	// 文字列の後の指定した文字列を削除する
	fmt.Println(strings.TrimRight("test", "t"))

	// Repeat
	// 文字列を繰り返す
	fmt.Println(strings.Repeat("test", 3))

	// Split
	// 文字列を分解する
	fmt.Println(strings.Split("test", "e"))

	// TimeSpace
	// 文字列の前後の空白を削除する
	fmt.Println(strings.TrimSpace(" test "))
}

// bufioパッケージ
func testBufio(){
	// 標準入力を行単位で読み込む
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// スキャンに失敗した場合のエラーを取得する
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー:", err)
	}
}


// ioutilパッケージ
func testIoutil(){
	// ioutil.ReadFile()
	// ファイルを読み込む
	// ioutil.WriteFile()
	// ファイルに書き込む
	// ioutil.TempFile()
	// 一時ファイルを作成する
	// ioutil.TempDir()
	// 一時ディレクトリを作成する
	// ioutil.ReadDir()
	// ディレクトリの内容を取得する
	// ioutil.NopCloser()
	// io.ReadCloserを返す
	// ioutil.Discard
	// io.Writerの破棄先



	//入力全体を読み込む
	// 大容量のファイルを読み込む場合は、一度に全てのデータを読み込むのではなく、一部ずつ読み込むことが推奨されている
	f, _ := os.Open("foo.txt")
	bs, _ := ioutil.ReadAll(f)
	fmt.Println(string(bs))

	// ファイルに書き込む
	if err := ioutil.WriteFile("bar.txt", bs, 0666); err != nil {
		log.Fatalln(err)
	}

}


func main() {
	testIoutil()
}