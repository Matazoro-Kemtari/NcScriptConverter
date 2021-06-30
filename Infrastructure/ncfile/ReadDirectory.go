package ncfile

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"nc-script-converter/Domain/ncscript"
)

type NcScriptDir struct{}

func NewNcScriptDir() ncscript.DirViewer {
	return new(NcScriptDir)
}

func (n *NcScriptDir) FetchDir(path string) ([]string, error) {
	if len(path) == 0 {
		return nil, fmt.Errorf("引数が空です")
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("ディレクトリ取得に失敗しました error:%v", err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			// ディレクトリは無視
			continue
		}
		paths = append(paths, filepath.Join(path, file.Name()))
	}

	return paths, nil
}
