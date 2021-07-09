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

	// NCファイルのフォルダ参照イベント
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
		}

		// ファイル一覧の有効化
		v.setListBoxesEnabled()

		// ファイル一覧の更新
		v.readFileListBox(p)

		// 全追加・削除ボタンの有効化
		v.setAllItemsButtonEnabled()

		// 順位ボタンの有効化
		v.setRankButtonEnabled()

		// 変換ボタンの有効化
		v.setConvertButtonEnabled()

		fmt.Println(v.inPath)
	})

	/* ファイル一覧選択時イベント */
	v.mainView.dirFilList.ConnectSelectionChanged(func(selected, deselected *core.QItemSelection) {
		// ファイル追加ボタンの有効化
		v.mainView.addButton.SetEnabled(v.mainView.dirFilList.CurrentItem().IsSelected())
	})

	/* ファイル一覧選択時イベント */
	v.mainView.inFileList.ConnectSelectionChanged(func(selected, deselected *core.QItemSelection) {
		// ファイル削除ボタンの有効化
		v.mainView.removeButton.SetEnabled(v.mainView.inFileList.CurrentItem().IsSelected())
	})

	// 全ファイル追加イベント
	v.mainView.allAddButton.ConnectClicked(func(checked bool) {
		v.mainView.inFileItems = append(v.mainView.inFileItems, v.mainView.dirFilItems...)
		v.mainView.dirFilItems = nil
		v.mainView.dirFilList.Clear()
		v.mainView.inFileList.Clear()
		v.mainView.inFileList.AddItems(v.mainView.inFileItems)

		v.setAllItemsButtonEnabled()
		v.setRankButtonEnabled()
		v.setConvertButtonEnabled()
	})

	// 全ファイル削除イベント
	v.mainView.allRemoveButton.ConnectClicked(func(checked bool) {
		v.mainView.dirFilItems = append(v.mainView.dirFilItems, v.mainView.inFileItems...)
		v.mainView.inFileItems = nil
		v.mainView.dirFilList.Clear()
		v.mainView.inFileList.Clear()
		v.mainView.dirFilList.AddItems(v.mainView.dirFilItems)

		v.setAllItemsButtonEnabled()
		v.setRankButtonEnabled()
		v.setConvertButtonEnabled()
	})

	// ファイル追加イベント
	v.mainView.addButton.ConnectClicked(func(checked bool) {
		name := v.mainView.dirFilList.CurrentItem().Text()
		v.mainView.inFileItems = append(v.mainView.inFileItems, name)
		v.mainView.inFileList.AddItem(name)

		i := v.mainView.dirFilList.CurrentIndex().Row()
		v.mainView.dirFilItems = remove(v.mainView.dirFilItems, i)
		v.mainView.dirFilList.Clear()
		v.mainView.dirFilList.AddItems(v.mainView.dirFilItems)

		v.setAllItemsButtonEnabled()
		v.setRankButtonEnabled()
		v.setConvertButtonEnabled()
	})

	// ファイル削除イベント
	v.mainView.removeButton.ConnectClicked(func(checked bool) {
		name := v.mainView.inFileList.CurrentItem().Text()
		v.mainView.dirFilItems = append(v.mainView.dirFilItems, name)
		v.mainView.dirFilList.AddItem(name)

		i := v.mainView.inFileList.CurrentIndex().Row()
		v.mainView.inFileItems = remove(v.mainView.inFileItems, i)
		v.mainView.inFileList.Clear()
		v.mainView.inFileList.AddItems(v.mainView.inFileItems)

		v.setAllItemsButtonEnabled()
		v.setRankButtonEnabled()
		v.setConvertButtonEnabled()
	})

	// 順位UPイベント
	v.mainView.raisableRankButton.ConnectClicked(func(checked bool) {
		i := v.mainView.inFileList.CurrentIndex().Row()
		if i <= 0 {
			return
		}
		buf := v.mainView.inFileItems[i-1]
		v.mainView.inFileItems[i-1] = v.mainView.inFileItems[i]
		v.mainView.inFileItems[i] = buf

		v.mainView.inFileList.Clear()
		v.mainView.inFileList.AddItems(v.mainView.inFileItems)
		v.mainView.inFileList.SetCurrentRow(i - 1)
	})

	// 順位DOWNイベント
	v.mainView.lowerableRankButton.ConnectClicked(func(checked bool) {
		i := v.mainView.inFileList.CurrentIndex().Row()
		if i < 0 || i == len(v.mainView.inFileItems)-1 {
			return
		}
		buf := v.mainView.inFileItems[i+1]
		v.mainView.inFileItems[i+1] = v.mainView.inFileItems[i]
		v.mainView.inFileItems[i] = buf

		v.mainView.inFileList.Clear()
		v.mainView.inFileList.AddItems(v.mainView.inFileItems)
		v.mainView.inFileList.SetCurrentRow(i + 1)
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
		}
		v.setConvertButtonEnabled()
		log.Println(v.outPath)
	})

	// コンバートの指定
	v.mainView.cnvButton.ConnectClicked(func(checked bool) {
		v.mainView.cnvButton.SetEnabled(!v.mainView.cnvButton.IsEnabled())
		// 起動
		if err := v.concat.ConcatenatedNcScript(v.inPath, v.mainView.inFileItems, v.outPath); err != nil {
			// エラーメッセージ
			log.Println("error:", "進捗リスト送信でエラー:", err)

			widgets.QMessageBox_Warning(
				v.mainView.window,
				"エラー情報",
				fmt.Sprintf(
					"深刻なエラーが発生しました:\n"+
						"進捗リストが送信できていない可能性があります\n"+
						"再実行しても回復しない場合は管理者へ連絡してください\n"+
						"%v",
					err,
				),
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

	v.setListBoxesEnabled()
	v.setConvertButtonEnabled()
	v.setAllItemsButtonEnabled()
	v.mainView.addButton.SetEnabled(false)
	v.mainView.removeButton.SetEnabled(false)
	v.setRankButtonEnabled()

	v.mainView.window.Show()
	widgets.QApplication_Exec()

}

/* 全追加・全削除の有効化 */
func (v *MainViewController) setAllItemsButtonEnabled() {
	v.mainView.allAddButton.SetEnabled(len(v.mainView.dirFilItems) > 0)
	v.mainView.allRemoveButton.SetEnabled(len(v.mainView.inFileItems) > 0)
}

/* 順位操作ボタンの有効化 */
func (v *MainViewController) setRankButtonEnabled() {
	v.mainView.raisableRankButton.SetEnabled(len(v.mainView.inFileItems) > 0)
	v.mainView.lowerableRankButton.SetEnabled(len(v.mainView.inFileItems) > 0)
}

/* 変換ボタンの有効化 */
func (v *MainViewController) setConvertButtonEnabled() {
	v.mainView.cnvButton.SetEnabled(len(v.mainView.inFileItems) > 0 && len(v.outPath) > 0)
}

/* ListBoxとボタンの有効化 */
func (v *MainViewController) setListBoxesEnabled() {
	var e bool
	if len(v.inPath) > 0 {
		if f, err := os.Stat(v.inPath); os.IsExist(err) || f.IsDir() {
			e = true
		}
	}

	v.mainView.dirFilList.SetEnabled(e)
	v.mainView.inFileList.SetEnabled(e)
}

/* ファイル一覧の更新 */
func (v *MainViewController) readFileListBox(p string) {
	if !v.concat.DirectoryExist(v.inPath) {
		return
	}

	// ListBoxのクリア
	v.mainView.dirFilList.Clear()
	v.mainView.dirFilItems = nil
	v.mainView.inFileList.Clear()
	v.mainView.inFileItems = nil
	v.mainView.addButton.SetEnabled(false)
	v.mainView.removeButton.SetEnabled(false)

	// ファイル一覧を取得する
	files, err := v.concat.FetchFileNames(p)
	if err != nil {
		return
	}
	v.mainView.dirFilItems = append(v.mainView.dirFilItems, files...)
	v.mainView.dirFilList.AddItems(v.mainView.dirFilItems)
}

func remove(s []string, i int) []string {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}
