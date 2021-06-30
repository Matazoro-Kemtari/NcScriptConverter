package ncfile

import (
	"fmt"
	"nc-script-converter/Domain/alterationncscript"
	"os"
)

type NcScriptFile struct{}

func NewNcScriptFile() alterationncscript.FileReader {
	return new(NcScriptFile)
}

func (n *NcScriptFile) ReadAll(path string) (*os.File, error) {
	if len(path) == 0 {
		return nil, fmt.Errorf("引数が空です")
	}

	fp, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("ファイルの読み込みに失敗しました error:%v", err)
	}

	return fp, nil
}
