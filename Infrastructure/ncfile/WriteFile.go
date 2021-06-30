package ncfile

import (
	"fmt"
	"nc-script-converter/Domain/alterationncscript"
	"os"
)

type WritableNcScriptFile struct{}

func NewWritableNcScriptFile() alterationncscript.FileWriter {
	return new(WritableNcScriptFile)
}

func (n *WritableNcScriptFile) WriteAll(path string, contents []string) error {
	if len(path) == 0 {
		return fmt.Errorf("引数 path が空です")
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0655)
	if err != nil {
		return fmt.Errorf("保存先がありません")
	}

	for i := range contents {
		if _, err := file.WriteString(contents[i] + "\r\n"); err != nil {
			return fmt.Errorf("%d行の保存に失敗しました error: %v", i, err)
		}
	}

	return nil
}
