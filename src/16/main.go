package main

import (
	"bufio"
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"text/template"
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

// regexpパッケージ
func testRegexp(){
	// regexp.Compile()
	// 正規表現をコンパイルする
	// regexp.MustCompile()
	// 正規表現をコンパイルする
	// regexp.Match()
	// 正規表現に一致するか判定する
	// regexp.MatchString()
	// 正規表現に一致するか判定する
	// regexp.Find()
	// 正規表現に一致する部分を取得する
	// regexp.FindAll()
	// 正規表現に一致する部分を全て取得する
	// regexp.FindString()
	// 正規表現に一致する部分を取得する
	// regexp.FindAllString()
	// 正規表現に一致する部分を全て取得する
	// regexp.ReplaceAll()
	// 正規表現に一致する部分を置換する
	// regexp.ReplaceAllString()
	// 正規表現に一致する部分を置換する
	// regexp.Split()
	// 正規表現に一致する部分で分割する
	// regexp.QuoteMeta()
	// メタ文字をエスケープする

	// Goの正規表現の基本の書き方
	// MatchString
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	// MatchStringは、正規表現に一致するか判定するが
	// 大容量のデータを扱うにはそのままでは使いにくい
	// 正規表現をコンパイルして使うことで、高速に処理することができる
	// Compile
	re1, _ := regexp.Compile("a([a-z]+)e")
	fmt.Println(re1.MatchString("apple"))

	// MustCompile
	// こちらの方がよく使われる
	re2 := regexp.MustCompile("a([a-z]+)e")
	fmt.Println(re2.MatchString("apple"))

	// エスケープ
	// メタ文字をエスケープする
	regexp.QuoteMeta("\\d")
	// バッククォートで囲むとバクスラは不要
	regexp.QuoteMeta(`a^b.c`)

	// 正規表現のフラグ
	// 正規表現のフラグを指定することで、正規表現の挙動を変更することができる
	// 正規表現のフラグは、正規表現の前に指定する
	// (?i)は、大文字小文字を区別しない
	// (?m)は、複数行モード
	// (?s)は、ドットが改行にもマッチする
	// (?U)は、非貪欲マッチ
	re3 := regexp.MustCompile(`(?i)apple`)
	fmt.Println(re3.MatchString("Apple"))

	// 幅を持たない正規表現
	// 正規表現には、幅を持たない正規表現と幅を持つ正規表現がある
	// 幅を持たない正規表現は、文字列の一部に一致する
	// ^ 文頭
	// $ 文末
	// \b 単語の境界
	// \B 単語の境界以外
	// \A 入力の先頭
	// \z 入力の末尾
	// \Z 入力の末尾または末尾の改行
	re4 := regexp.MustCompile(`^apple$`) //使用頻度が高い！
	fmt.Println(re4.MatchString("apple")) //true
	fmt.Println(re4.MatchString("   apple")) //false
	fmt.Println(re4.MatchString("apple   ")) //false

	// 繰り返しを表す正規表現
	// 正規表現には、繰り返しを表す正規表現がある
	// * 0回以上の繰り返し
	// + 1回以上の繰り返し
	// ? 0回または1回の繰り返し
	// {n} n回の繰り返し
	// {n,} n回以上の繰り返し
	// {n,m} n回以上m回以下の繰り返し
	re5 := regexp.MustCompile(`a+`)
	fmt.Println(re5.MatchString("a")) //true
	fmt.Println(re5.MatchString("aa")) //true
	fmt.Println(re5.MatchString("aaa")) //true
	fmt.Println(re5.MatchString("b")) //false

	// 正規表現の文字クラス
	// 正規表現には、文字クラスを表す正規表現がある
	// [] 文字クラス
	// [a-z] aからzまでの文字
	// [^a-z] aからz以外の文字
	// [a-zA-Z] aからzとAからZまでの文字
	// [0-9] 0から9までの数字
	// [^0-9] 0から9以外の文字
	// [a-z0-9] aからzと0から9までの文字
	re6 := regexp.MustCompile(`[a-z]+`)
	fmt.Println(re6.MatchString("apple")) //true
	fmt.Println("///")
	fmt.Println(re6.MatchString("Apple")) //true
	fmt.Println(re6.MatchString("123")) //false

	// 正規表現のグループ
	// 正規表現には、グループを表す正規表現がある
	// () グループ
	// (a|b) aまたはb
	// (?:a) キャプチャしないグループ
	// (?P<name>a) 名前付きグループ
	re7 := regexp.MustCompile(`(a|A)(z|Z)`)
	fmt.Println(re7.MatchString("az")) //true
	fmt.Println(re7.MatchString("aZ")) //true
	fmt.Println(re7.MatchString("Az")) //true
	fmt.Println(re7.MatchString("AA")) //false

	//FindString
	// 正規表現に一致する部分を取得する
	re8 := regexp.MustCompile(`(a|A)(z|Z)`)
	fmt.Println(re8.FindString("ABCXYZ")) //nil
	fmt.Println(re8.FindString("ABCXYZAZAZAZ")) //AZ
	fmt.Println(re8.FindString("banana")) //nil
	//FindAllString
	// 正規表現に一致する部分を全て取得する
	//代に引数に-1を指定すると、全ての一致する部分を取得する
	fmt.Println(re8.FindAllString("ABCXYZAZAZAZ", -1)) //[AZ AZ AZ]

	// 正規表現に一致する部分を取得する
	// FindStringSubmatch
	// 正規表現に一致する部分を全て取得する
	// FindAllStringSubmatch
	// 第二引数に-1を指定すると、全ての一致する部分を取得する
	re9 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	fmt.Println(re9.FindStringSubmatch("2024-10-13")) //[2024-10-13 2024 10 13]
	//[[2024-10-13 2024 10 13] [2024-10-12 2024 10 12] [2024-10-15 2024 10 15]]
	fmt.Println(re9.FindAllStringSubmatch("2024-10-13 2024-10-12 2024-10-15", -1))

	// 正規表現に一致する部分を置換する
	// ReplaceAllString
	re10 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	fmt.Println(re10.ReplaceAllString("2024-10-13", "$3/$2/$1")) //13/10/2024

	// 正規表現に一致する部分で分割する
	// Split
	re11 := regexp.MustCompile(`\s+`)
	fmt.Println(re11.Split("a b c", -1)) //[a b c]
}

// syncパッケージ
func testSync(){
	// レースコンディション
	// 複数のゴルーチンが同時に同じ変数にアクセスすることで、予期しない結果が発生すること
	// レースコンディションを避けるためには、排他制御を行う必要がある
	// sync.Mutex
	// 排他制御を行う
	// sync.RWMutex
	// 読み込みと書き込みの排他制御を行う
	// sync.WaitGroup
	// ゴルーチンの完了を待つ
	// sync.Once
	// 一度だけ実行する
	// sync.Cond
	// ゴルーチン間での通信を行う

	var mutex *sync.Mutex
	mutex.Lock()
	mutex.Unlock()

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("A")
	}()
	wg.Wait()
}

// criptoパッケージ
func testCrypto(){
	// cryptoパッケージ
	// 暗号化と復号化を行う
	// crypto/aes
	// crypto/cipher
	// crypto/des
	// crypto/dsa
	// crypto/ecdsa
	// crypto/elliptic
	// crypto/hmac
	// crypto/md5
	// crypto/rand
	// crypto/rc4
	// crypto/rsa
	// crypto/sha1
	// crypto/sha256
	// crypto/sha512
	// crypto/subtle
	// crypto/tls
	// crypto/x509

	//MD5ハッシュ値を計算する
	h:=md5.New()
	io.WriteString(h, "test")
	fmt.Println(h.Sum(nil))

	s:= fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s)


	    //------------------------------
    /* SHA-1 */
    s1 := sha1.New()
    io.WriteString(s1, "ABCDE")
    fmt.Printf("%x\n", s1.Sum(nil))
    // => f0393febe8baaa55e32f7be2a7cc180bf34e52137d99e056c817a9c07b8f239a

    //------------------------------

    /* SHA-256 */
    s256 := sha256.New()
    io.WriteString(s256, "ABCDE")
    fmt.Printf("%x\n", s256.Sum(nil))
    // => f0393febe8baaa55e32f7be2a7cc180bf34e52137d99e056c817a9c07b8f239a

    //------------------------------

    /* SHA-512 */
    s512 := sha512.New()
    io.WriteString(s512, "ABCDE")
    fmt.Printf("%x\n", s512.Sum(nil))
	// => 9989a8fcbc29044b5883a0a36c146fe7415b1439e995b4d806ea0af7da9ca4390eb92a604b3ecfa3d75f9911c768fbe2aecc59eff1e48dcaeca1957bdde01dfb
}

// Jsonパッケージ
func testJson(){
	// jsonパッケージ
	// JSON形式のデータを扱う
	// json.Marshal()
	// GoのデータをJSON形式に変換する
	// json.Unmarshal()
	// JSON形式のデータをGoのデータに変換する
	// json.NewEncoder()
	// JSON形式のデータをエンコードする
	// json.NewDecoder()
	// JSON形式のデータをデコードする

	u :=new(User)
	u.ID = 1
	u.Name = "test"
	u.Age = 20
	u.Email = "email.com"
	// u.A = A{}
	// Json形式に変換する
	// Marshal
	// byteのスライスに変換する
	bs, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	u2 := new(User)
	// JSON形式のデータをGoのデータに変換する
	// Unmarshal
	if err := json.Unmarshal(bs, u2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(u2)

	// MarchalJSON
	// JSON形式に変換する時に編集できる
	// UnmarchalJSON
	// JSON形式のデータをStructに変換する時に編集できる

}
//jsonに変換した時の形式を指定する
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name,string"` //jsonに変換する時にstringに変換する
	Age  int    `json:"-"` //jsonに変換しない
	Email string `json:"email"`
	A   *A      `json:"a,omitempty"`//値が空の場合は出力しない
}
type A struct {}

// sortパッケージ
func testSort(){
	// sortパッケージ
	// スライスをソートする

	// sort.Ints()
	// int型のスライスをソートする

	// sort.Float64s()
	// float64型のスライスをソートする

	// sort.Strings()
	// string型のスライスをソートする

	// sort.Slice()
	// スライスをソートする

	// sort.Search()
	// ソートされたスライスから値を検索する

	// sort.SearchInts()
	// ソートされたint型のスライスから値を検索する

	// sort.SearchFloat64s()
	// ソートされたfloat64型のスライスから値を検索する

	// sort.SearchStrings()
	// ソートされたstring型のスライスから値を検索する

	// sort.SearchSlice()
	// ソートされたスライスから値を検索する

	// sort.Reverse()
	// スライスを逆順にする

	// sort.IsSorted()
	// スライスがソートされているか判定する

	// sort.Stable()
	// 安定ソートを行う

	// sort.Interface
	type Enty struct {
		Name string
		Age  int
	}

	el := []Enty{{"A", 3}, {"F", 1}, {"i", 4}, {"b",1}, {"t", 5}, {"y",9}, {"c", 2}}
	sort.Slice(el, func(i, j int) bool {
		return el[i].Name < el[j].Name
	});
	fmt.Println(el)

	sort.SliceStable(el, func(i, j int) bool {
		return el[i].Name < el[j].Name
	});
	fmt.Println(el)
}

// contextパッケージ
func testContext(){
	// contextパッケージ
	// ゴルーチン間で値を共有する
	// context.Context
	// ゴルーチン間で値を共有する
	// context.Background()
	// 空のContextを作成する
	// context.TODO()
	// 未実装のContextを作成する
	// context.WithValue()
	// Contextに値を設定する
	// context.WithCancel()
	// Contextをキャンセルする
	// context.WithDeadline()
	// Contextに期限を設定する
	// context.WithTimeout()
	// Contextにタイムアウト

	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	go func() {
		fmt.Println("start")
		time.Sleep(2 * time.Second)
		fmt.Println("end")
		ch <- "result"
	}()
}

// net/urlパッケージ
func testUrl(){
	// net/urlパッケージ
	// URLをパースする
	// url.Parse()
	// URLをパースする
	// url.ParseRequestURI()
	// URLをパースする
	// url.ParseQuery()
	// クエリをパースする
	// url.QueryEscape()
	// クエリをエスケープする
	// url.QueryUnescape()
	// クエリをアンエスケープする

	// URLをパースする
	u, _ := url.Parse("https://example.com:8080/path?key=value#fragment")
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)
	fmt.Println(u.Fragment)
	// https
	// example.com:8080
	// /path
	// key=value
	// fragment

	// URLを生成する
	u2 := &url.URL{
		Scheme:   "https",
		Host:     "example.com:8080",
	}
	q:=u2.Query()
	q.Set("q", "golang")
	u2.RawQuery = q.Encode()
	// https://example.com:8080?q=golang
	fmt.Println(u2)
}

// net/httpパッケージ
// クライアント
func testHttpClient(){
	// net/httpパッケージ
	// HTTPクライアントを作成する
	// http.Get()
	// GETリクエストを送信する
	// http.Post()
	// POSTリクエストを送信する
	// http.PostForm()
	// POSTリクエストを送信する
	// http.NewRequest()
	// リクエストを作成する
	// http.Client
	// カスタムのHTTPクライアントを作成する

	// GETリクエストを送信する
	resp, _ := http.Get("https://example.com")
	fmt.Println(resp.StatusCode)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	// POSTリクエストを送信する
	resp, err := http.PostForm("https://example.com", url.Values{"key": {"value"}})
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	// リクエストを作成する
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	req.Header.Set("Content-Type", "text/plain")
	client := &http.Client{}
	resp, _ = client.Do(req)
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	//応用
	// //ヘッダーをつけたり、クエリをつけたり
	// //Parse 正しいURLか確認
	// base, _ := url.Parse("https://example.com/")
 
	// //クエリ の確認 URLの後につく
	// reference, _ := url.Parse("index/lists?id=1")
 
	// //ResolveReference
	// //クエリをくっつけたURLを生成する。
	// //相対パスから絶対URLに変換する。
	// //baseのURLの末尾に文字列が入っていたとしても、正しいURLに直してくれる
	// endpoint := base.ResolveReference(reference).String()
	// fmt.Println(endpoint)
 
	// //GET ver
	// //リクエストを作成 nil部はPOST時のみ設定（バイトを入れる）
	// //まだリクエストはしていない。structを作っただけ。
	// req, _ := http.NewRequest("GET", endpoint, nil)
 
	// //requestにheaderをつける。cash情報など
	// req.Header.Add("Content-Type", `application/json"`)
 
	// //URLのクエリを確認
	// q := req.URL.Query()
 
	// //クエリを追加
	// q.Add("name", "test")
 
	// //クエリを表示
	// fmt.Println(q)
 
	// //&など特殊文字などがある場合があるため、encodingしてからURLに追加してやる必要がある
	// fmt.Println(q.Encode())
 
	// //encodeしてからURLに戻す
	// //日本語などを変換する
	// req.URL.RawQuery = q.Encode()
 
	// //実際にアクセスする
	// //クライアントを作る
	// var client *http.Client = &http.Client{}
 
	// //結果 アクセスする
	// resp, _ := client.Do(req)
 
	// //読み込み
	// body, _ := ioutil.ReadAll(resp.Body)
 
	// //出力
	// fmt.Println(string(body))

	// Post ver
	//応用
    // //ヘッダーをつけたり、クエリをつけたり
    // //Parse 正しいURLか確認
    // base, _ := url.Parse("https://example.com/")
 
    // //クエリ の確認 URLの後につく
    // reference, _ := url.Parse("index/lists?id=1")
 
    // //Postの時のデータ
    // AccountData := &Account{ID: "ABC-DEF", PassWord: "testtest"}
    // data, _ := json.Marshal(AccountData)
 
    // //ResolveReference
    // //クエリをくっつけたURLを生成する。
    // //相対パスから絶対URLに変換する。
    // //baseのURLの末尾に文字列が入っていたとしても、正しいURLに直してくれる
    // endpoint := base.ResolveReference(reference).String()
    // fmt.Println(endpoint)
 
    // //POST ver
    // //bytes.NewBuffer([]byte("password")でリクエストの領域を作成
    // //POSTの場合は、Bodyにデータを入れる。例えばパスワード。見られたらダメな情報はbodyに
    // req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
 
    // //requestにheaderをつける。cash情報など
    // req.Header.Add("Content-Type", `application/json"`)
 
    // //URLのクエリを確認
    // q := req.URL.Query()
 
    // //クエリを追加
    // q.Add("name", "test")
 
    // //クエリを表示
    // fmt.Println(q)
 
    // //&など特殊文字などがある場合があるため、encodingしてからURLに追加してやる必要がある
    // fmt.Println(q.Encode())
 
    // //encodeしてからURLに戻す
    // //日本語などを変換する
    // req.URL.RawQuery = q.Encode()
 
    // //実際にアクセスする
    // //クライアントを作る
    // var client *http.Client = &http.Client{}
 
    // //結果 アクセスする
    // resp, _ := client.Do(req)
 
    // //読み込み
    // body, _ := ioutil.ReadAll(resp.Body)
 
    // //出力
    // fmt.Println(string(body))
}

// net/http server
// type MyHandler struct{}
// func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, World")
// }
func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, "Hello, World111")
}

func testServer(){
	http.HandleFunc("/top", top)
	http.ListenAndServe(":8080", nil)
}

func main() {
	testServer()
}