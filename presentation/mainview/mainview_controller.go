package mainview

import (
	"fmt"
	"log"
	"nc-script-converter/UseCase/concatenatedscript"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type MainViewController struct {
	concat    *concatenatedscript.ConcatenatedNcScriptUseCase
	inPath    string
	outPath   string
	inLabel   *widgets.QLabel
	outLabel  *widgets.QLabel
	cnvButton *widgets.QPushButton
}

func NewMainViewController(concat *concatenatedscript.ConcatenatedNcScriptUseCase) *MainViewController {
	return &MainViewController{
		concat: concat,
	}
}

func (v *MainViewController) Initialize() {
	// fmt.Println("[NCデータが保存されているフォルダを指定してください]")
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// in := scanner.Text()

	// fmt.Println("[結合ファイルのファイル名を指定してください]")
	// scanner.Scan()
	// out := scanner.Text()
	// (*v.concat).ConcatenatedNcScript(in, out)

	// fmt.Println("[終了しました。何かキーを押して終了します]")
	// scanner.Scan()

	v.showWindow()
}

func (v *MainViewController) showWindow() {

	// 参考
	// https://saitodev.co/article/Go%E8%A8%80%E8%AA%9E%E3%81%A7Qt%E3%81%AEQFormLayout%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B/
	// https://saitodev.co/article/Go%E8%A8%80%E8%AA%9E%E3%81%A7Qt%E3%81%AEQBoxLayout%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B

	// Base Window 作成
	core.QCoreApplication_SetApplicationName("NcScriptConverter")
	core.QCoreApplication_SetOrganizationName("company")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(400, 300)
	window.SetWindowTitle("NCプログラム変換")

	//フレームワークに上記で作成したレイアウトをセットする
	baseWidget := widgets.NewQWidget(nil, 0)
	//フレームワークにQTBoxLayoutをはめ込む
	//第一引数で0は左から右、1は右から左、2は上から下、3は下から上
	vbox := widgets.NewQBoxLayout(2, nil)
	baseWidget.SetLayout(vbox)

	// NCデータフォルダの指定
	hbox1 := widgets.NewQHBoxLayout()
	label1 := widgets.NewQLabel2("NCデータフォルダ", nil, 0)
	inButton := widgets.NewQPushButton2("参照", nil)
	hbox1.AddWidget(label1, 0, core.Qt__AlignTrailing)
	hbox1.AddWidget(inButton, 0, core.Qt__AlignBaseline)
	(*v).inLabel = widgets.NewQLabel2("", nil, 0)

	// フォルダ参照イベント
	inButton.ConnectClicked(func(checked bool) {
		// フォルダ選択ダイアログ表示
		// https://day-journal.com/memo/qt-005/
		// http://qt-log.open-memo.net/sub/dialog__directory_dialog.html
		p := widgets.QFileDialog_GetExistingDirectory(
			nil,
			"NCデータフォルダを指定してください",
			"C:\\",
			widgets.QFileDialog__DontUseCustomDirectoryIcons,
		)
		if len(p) > 0 {
			(*v).inPath = p
			v.inLabel.SetText(p)
			v.setConvertButtonEnabled()
		}
		fmt.Println(v.inPath)
	})

	// 結合ファイルの保存先
	hbox2 := widgets.NewQHBoxLayout()
	label2 := widgets.NewQLabel2("結合ファイルの保存先", nil, 0)
	outButton := widgets.NewQPushButton2("参照", nil)
	hbox2.AddWidget(label2, 0, core.Qt__AlignTrailing)
	hbox2.AddWidget(outButton, 0, core.Qt__AlignBaseline)
	(*v).outLabel = widgets.NewQLabel2("", nil, 0)

	outButton.ConnectClicked(func(checked bool) {
		// ファイルの保存先ダイアログ表示
		p := widgets.QFileDialog_GetSaveFileName(
			nil,
			"結合ファイル名を指定してください",
			"C:\\",
			"すべて(*.*)",
			"すべて(*.*)",
			widgets.QFileDialog__DontUseCustomDirectoryIcons,
		)
		if len(p) > 0 {
			(*v).outPath = p
			v.outLabel.SetText(p)
			v.setConvertButtonEnabled()
		}
		log.Println(v.outPath)
	})

	// コンバートの指定
	v.cnvButton = widgets.NewQPushButton2("実行", nil)
	v.cnvButton.ConnectClicked(func(checked bool) {
		v.cnvButton.SetEnabled(!v.cnvButton.IsEnabled())
		// 起動
		if err := v.concat.ConcatenatedNcScript(v.inPath, v.outPath); err != nil {
			// エラーメッセージ
			log.Println("error:", "進捗リスト送信でエラー:", err)

			widgets.QMessageBox_Warning(
				window,
				"エラー情報",
				fmt.Sprintf("深刻なエラーが発生しました:\n"+
					"進捗リストが送信できていない可能性があります\n"+
					"再実行しても回復しない場合は管理者へ連絡してください\n"+
					"%v", err),
				widgets.QMessageBox__Ok,
				widgets.QMessageBox__Ok,
			)

			inButton.SetEnabled(!inButton.IsEnabled())
			return
		}

		// 終了メッセージ
		widgets.QMessageBox_Information(
			window,
			"正常終了",
			"結合処理が正常に終了しました",
			widgets.QMessageBox__Ok,
			widgets.QMessageBox__Ok,
		)
		v.inLabel.Clear()
		v.outLabel.Clear()
		// v.cnvButton.SetEnabled(!v.cnvButton.IsEnabled())
	})

	v.setConvertButtonEnabled()

	window.SetCentralWidget(baseWidget)
	vbox.AddLayout(hbox1, 0)
	vbox.AddWidget(v.inLabel, 0, core.Qt__AlignBaseline)
	vbox.AddLayout(hbox2, 0)
	vbox.AddWidget(v.outLabel, 0, core.Qt__AlignBaseline)
	vbox.AddWidget(v.cnvButton, 0, core.Qt__AlignBaseline)
	window.Show()

	widgets.QApplication_Exec()

}

func (v *MainViewController) setConvertButtonEnabled() {
	v.cnvButton.SetEnabled(len(v.inPath) > 0 && len(v.outPath) > 0)
}
