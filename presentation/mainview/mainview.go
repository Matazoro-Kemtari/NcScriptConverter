package mainview

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

/*
QListViewのItemを消したりするのができない
https://github.com/therecipe/examples/blob/master/advanced/widgets/listview/main.go
だからSliceで代用
*/
type MainView struct {
	version             string
	window              *widgets.QMainWindow
	inLabel             *widgets.QLabel
	inButton            *widgets.QPushButton
	dirFilList          *widgets.QListWidget
	dirFilItems         []string
	allAddButton        *widgets.QPushButton
	addButton           *widgets.QPushButton
	removeButton        *widgets.QPushButton
	allRemoveButton     *widgets.QPushButton
	inFileList          *widgets.QListWidget
	inFileItems         []string
	raisableRankButton  *widgets.QPushButton
	lowerableRankButton *widgets.QPushButton
	outButton           *widgets.QPushButton
	outLabel            *widgets.QLabel
	openReviewCheckBox  *widgets.QCheckBox
	cnvButton           *widgets.QPushButton
}

func NewMainView(version string) *MainView {
	mv := &MainView{
		version: version,
	}
	mv.Initialize()
	return mv
}

func (m *MainView) Initialize() {
	m.createWindow()
	m.viewDesign()
	m.addStyle()
}

func (m *MainView) createWindow() {

	// 参考
	// https://saitodev.co/article/Go%E8%A8%80%E8%AA%9E%E3%81%A7Qt%E3%81%AEQFormLayout%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B/
	// https://saitodev.co/article/Go%E8%A8%80%E8%AA%9E%E3%81%A7Qt%E3%81%AEQBoxLayout%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B
	// https://qiita.com/shu32/items/53204832c1074ff7cd7f

	// Base Window 作成
	core.QCoreApplication_SetApplicationName("NcScriptConverter")
	core.QCoreApplication_SetOrganizationName("company")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	widgets.NewQApplication(len(os.Args), os.Args)

	m.window = widgets.NewQMainWindow(nil, 0)
	m.window.SetMinimumSize2(800, 400)
	m.window.SetWindowTitle("NCプログラム変換 : " + m.version)

}

func (m *MainView) viewDesign() {

	//フレームワークに上記で作成したレイアウトをセットする
	baseWidget := widgets.NewQWidget(nil, 0)
	//フレームワークにQTBoxLayoutをはめ込む
	//第一引数で0は左から右、1は右から左、2は上から下、3は下から上
	vbox := widgets.NewQBoxLayout(2, nil)
	baseWidget.SetLayout(vbox)

	// NCデータフォルダの指定
	label1 := widgets.NewQLabel2("NCデータフォルダ", nil, 0)
	(*m).inLabel = widgets.NewQLabel2("", nil, 0)
	// 効果が見られなかった
	// http://qt-log.open-memo.net/sub/layout__how_to_use_size_policy.html
	// m.inLabel.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Fixed)
	(*m).inButton = widgets.NewQPushButton2("参照", nil)
	hbox1 := widgets.NewQHBoxLayout()
	hbox1.SetSpacing(5)
	hbox1.AddWidget(label1, 0, core.Qt__AlignLeft)
	hbox1.AddWidget((*m).inLabel, 0, core.Qt__AlignBaseline)
	hbox1.AddWidget((*m).inButton, 0, core.Qt__AlignRight)

	// ディレクトリ内のファイル一覧
	label3 := widgets.NewQLabel2("ファイル一覧", nil, 0)
	(*m).dirFilList = widgets.NewQListWidget(nil)
	vbox1 := widgets.NewQVBoxLayout()
	vbox1.AddWidget(label3, 0, core.Qt__AlignBaseline)
	vbox1.AddWidget(m.dirFilList, 0, core.Qt__AlignBaseline)

	// 処理対象操作
	(*m).allAddButton = widgets.NewQPushButton2("→ →", nil)
	(*m).addButton = widgets.NewQPushButton2("→", nil)
	(*m).removeButton = widgets.NewQPushButton2("←", nil)
	(*m).allRemoveButton = widgets.NewQPushButton2("← ←", nil)
	vbox2 := widgets.NewQVBoxLayout()
	vbox2.AddWidget(m.allAddButton, 0, core.Qt__AlignBaseline)
	vbox2.AddWidget(m.addButton, 0, core.Qt__AlignBaseline)
	vbox2.AddWidget(m.removeButton, 0, core.Qt__AlignBaseline)
	vbox2.AddWidget(m.allRemoveButton, 0, core.Qt__AlignBaseline)

	// 対象NCファイル一覧
	label4 := widgets.NewQLabel2("対象NCファイル", nil, 0)
	(*m).inFileList = widgets.NewQListWidget(nil)
	vbox3 := widgets.NewQVBoxLayout()
	vbox3.AddWidget(label4, 0, core.Qt__AlignBaseline)
	vbox3.AddWidget(m.inFileList, 0, core.Qt__AlignBaseline)

	// 並び替え操作
	vbox4 := widgets.NewQVBoxLayout()
	(*m).raisableRankButton = widgets.NewQPushButton2("[↑]", nil)
	(*m).lowerableRankButton = widgets.NewQPushButton2("[↓]", nil)
	vbox4.AddWidget(m.raisableRankButton, 0, core.Qt__AlignBaseline)
	vbox4.AddWidget(m.lowerableRankButton, 0, core.Qt__AlignBaseline)
	hbox3 := widgets.NewQHBoxLayout()

	// 処理対象対象選択
	hbox3.AddLayout(vbox1, 0)
	hbox3.AddLayout(vbox2, 0)
	hbox3.AddLayout(vbox3, 0)
	hbox3.AddLayout(vbox4, 0)

	// 結合ファイルの保存先
	label2 := widgets.NewQLabel2("結合ファイルの保存先", nil, 0)
	(*m).outLabel = widgets.NewQLabel2("", nil, 0)
	(*m).outButton = widgets.NewQPushButton2("参照", nil)
	hbox2 := widgets.NewQHBoxLayout()
	hbox2.AddWidget(label2, 0, core.Qt__AlignLeft)
	hbox2.AddWidget(m.outLabel, 0, core.Qt__AlignBaseline)
	hbox2.AddWidget(m.outButton, 0, core.Qt__AlignRight)

	// オープンレビューの選択
	m.openReviewCheckBox = widgets.NewQCheckBox2("オープンレビューで実行する", nil)

	// コンバートの指定
	m.cnvButton = widgets.NewQPushButton2("実行", nil)
	m.cnvButton.SetMinimumHeight(50)
	m.cnvButton.SetMaximumHeight(70)

	m.window.SetCentralWidget(baseWidget)
	vbox.AddLayout(hbox1, 0)
	// vbox.AddWidget(v.inLabel, 0, core.Qt__AlignBaseline)
	vbox.AddLayout(hbox3, 0)
	vbox.AddLayout(hbox2, 0)
	vbox.AddWidget(m.openReviewCheckBox, 0, core.Qt__AlignBaseline)
	vbox.AddWidget(m.cnvButton, 0, core.Qt__AlignBaseline)

}

func (m *MainView) addStyle() {
	// (*m).inLabel.SetStyleSheet(
	// 	"border:1px solid black;",
	// )
}
