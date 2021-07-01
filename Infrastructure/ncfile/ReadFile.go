package ncfile

import (
	"bufio"
	"fmt"
	"nc-script-converter/Domain/alterationncscript"
	"os"
)

type ReadableNcScriptFile struct{}

func NewReadableNcScriptFile() *alterationncscript.FileReader {
	var obj alterationncscript.FileReader = &ReadableNcScriptFile{}
	return &obj
}

func (n *ReadableNcScriptFile) ReadAll(path string) ([]string, error) {
	if len(path) == 0 {
		return nil, fmt.Errorf("引数が空です")
	}

	fp, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("ファイルの読み込みに失敗しました error:%v", err)
	}
	defer fp.Close()

	s := bufio.NewScanner(fp)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines, nil
}
