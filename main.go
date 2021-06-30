package main

import (
	"flag"
	"fmt"
	"log"
	"nc-script-converter/presentation/mainview"
	"os"
	"time"

	"github.com/comail/colog"
)

// バージョン埋め込む
// INFO: https://qiita.com/irotoris/items/4aae9ad483bf08915688
var version string
var revision string

func main() {
	// コマンドライン引数を取得
	isVersion := flag.Bool("version", false, "バージョンを表示する")
	flag.Parse()

	// バージョン表示
	if *isVersion {
		fmt.Printf("version: %s-%s\n", version, revision)
		os.Exit(0)
	}

	// logの設定 https://qiita.com/kmtr/items/406073651d7a12aab9c6
	file, err := os.OpenFile("error_json.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0655)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	colog.SetOutput(file)
	colog.ParseFields(true)
	colog.SetFormatter(&colog.JSONFormatter{
		TimeFormat: time.RFC3339,
		Flag:       log.Lshortfile,
	})
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	// colog.SetFormatter(&colog.StdFormatter{
	// 	Colors: true,
	// 	Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	// })
	colog.Register()

	mainview.ShowWindow()
}
