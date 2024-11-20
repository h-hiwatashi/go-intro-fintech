package utils

import (
	"io"
	"log"
	"os"
)

func logingSettings(logFile string){
	// 06666 is the permission for the file
	// logfileを読み込む
	// ファイルがない場合は作成
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	// MultiWriterで標準出力とlogfileに出力
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// logの設定
	// 日付、時間、ファイル名を出力
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// 出力先を設定
	log.SetOutput(multiLogFile)
}