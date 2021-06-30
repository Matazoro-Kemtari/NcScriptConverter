package mainview

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func ShowWindow() {

	// 参考
	// https://saitodev.co/article/Go%E8%A8%80%E8%AA%9E%E3%81%A7Qt%E3%81%AEQFormLayout%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B/
	// https://saitodev.co/article/Go%E8%A8%80%E8%AA%9E%E3%81%A7Qt%E3%81%AEQBoxLayout%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B
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

	hbox := widgets.NewQHBoxLayout()
	label1 := widgets.NewQLabel2("NCデータフォルダ", nil, 0)
	fButton := widgets.NewQPushButton2("参照", nil)
	hbox.AddWidget(label1, 0, core.Qt__AlignCenter)
	hbox.AddWidget(fButton, 0, core.Qt__AlignCenter)
	fButton.ConnectClicked(func(checked bool) {
		// フォルダ選択ダイアログ表示
		// https://day-journal.com/memo/qt-005/
		// http://qt-log.open-memo.net/sub/dialog__directory_dialog.html
		dir := widgets.QFileDialog_GetExistingDirectory(
			nil,
			"NCデータフォルダを指定してください",
			"C:\\",
			widgets.QFileDialog__DontConfirmOverwrite,
		)
		fmt.Println(dir)

	})
	cnvButton := widgets.NewQPushButton2("実行", nil)
	cnvButton.ConnectClicked(func(checked bool) {
		cnvButton.SetEnabled(!cnvButton.IsEnabled())
		// 起動
		// if err := c.ExecuteShareDocument(); err != nil {
		// 	// エラーメッセージ
		// 	log.Println("error:", "進捗リスト送信でエラー:", err)

		// 	widgets.QMessageBox_Warning(
		// 		window,
		// 		"エラー情報",
		// 		fmt.Sprintf("深刻なエラーが発生しました:\n"+
		// 			"進捗リストが送信できていない可能性があります\n"+
		// 			"再実行しても回復しない場合は管理者へ連絡してください\n"+
		// 			"%v", err),
		// 		widgets.QMessageBox__Ok,
		// 		widgets.QMessageBox__Ok,
		// 	)

		// 	button.SetEnabled(!button.IsEnabled())
		// 	return
		// }
		cnvButton.SetEnabled(!cnvButton.IsEnabled())
	})

	// baseWidget.SetLayout(formLayout)
	window.SetCentralWidget(baseWidget)
	baseWidget.Layout().AddChildLayout(hbox)
	// baseWidget.Layout().AddWidget(fButton)
	baseWidget.Layout().AddWidget(cnvButton)
	window.Show()

	widgets.QApplication_Exec()

}
