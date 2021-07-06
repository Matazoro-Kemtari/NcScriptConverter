package mainview

import (
	"fmt"
	"log"
	"nc-script-converter/UseCase/concatenatedscript"

	"github.com/therecipe/qt/widgets"
)

type MainViewController struct {
	version  string
	concat   *concatenatedscript.ConcatenatedNcScriptUseCase
	inPath   string
	outPath  string
	mainView *MainView
}

func NewMainViewController(version string, concat *concatenatedscript.ConcatenatedNcScriptUseCase) *MainViewController {
	return &MainViewController{
		version:  version,
		concat:   concat,
		mainView: NewMainView(version),
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

	// フォルダ参照イベント
	v.mainView.inButton.ConnectClicked(func(checked bool) {
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
			v.mainView.inLabel.SetText(p)
			v.setConvertButtonEnabled()
		}
		fmt.Println(v.inPath)
	})

	// 結合ファイルの保存先参照イベント
	v.mainView.outButton.ConnectClicked(func(checked bool) {
		// ファイルの保存先ダイアログ表示
		p := widgets.QFileDialog_GetSaveFileName(
			nil,
			"結合ファイル名を指定してください",
			"C:\\",
			"すべて(*)",
			"すべて(*)",
			widgets.QFileDialog__DontUseCustomDirectoryIcons,
		)
		if len(p) > 0 {
			(*v).outPath = p
			v.mainView.outLabel.SetText(p)
			v.setConvertButtonEnabled()
		}
		log.Println(v.outPath)
	})

	// コンバートの指定
	v.mainView.cnvButton.ConnectClicked(func(checked bool) {
		v.mainView.cnvButton.SetEnabled(!v.mainView.cnvButton.IsEnabled())
		// 起動
		if err := v.concat.ConcatenatedNcScript(v.inPath, v.outPath); err != nil {
			// エラーメッセージ
			log.Println("error:", "進捗リスト送信でエラー:", err)

			widgets.QMessageBox_Warning(
				v.mainView.window,
				"エラー情報",
				fmt.Sprintf("深刻なエラーが発生しました:\n"+
					"進捗リストが送信できていない可能性があります\n"+
					"再実行しても回復しない場合は管理者へ連絡してください\n"+
					"%v", err),
				widgets.QMessageBox__Ok,
				widgets.QMessageBox__Ok,
			)

			v.mainView.inButton.SetEnabled(!v.mainView.inButton.IsEnabled())
			return
		}

		// 終了メッセージ
		widgets.QMessageBox_Information(
			v.mainView.window,
			"正常終了",
			"結合処理が正常に終了しました",
			widgets.QMessageBox__Ok,
			widgets.QMessageBox__Ok,
		)
		v.mainView.inLabel.Clear()
		v.mainView.outLabel.Clear()
		// v.cnvButton.SetEnabled(!v.cnvButton.IsEnabled())
	})

	v.setConvertButtonEnabled()

	v.mainView.window.Show()
	widgets.QApplication_Exec()

}

func (v *MainViewController) setConvertButtonEnabled() {
	v.mainView.cnvButton.SetEnabled(len(v.inPath) > 0 && len(v.outPath) > 0)
}
