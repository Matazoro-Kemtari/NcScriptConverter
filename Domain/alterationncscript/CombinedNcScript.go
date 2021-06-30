package alterationncscript

import (
	"bufio"
	"fmt"
)

type CombinedNcScript struct {
	dir DirViewer
	fr  FileReader
}

func NewCombinedNcScript(dir DirViewer, fr FileReader) *CombinedNcScript {
	return &CombinedNcScript{
		dir: dir,
		fr:  fr,
	}
}

func (c *CombinedNcScript) CombineNcScript(inPath string, outPath string) error {
	// ファイル一覧取得
	files, err := c.dir.FetchDir(inPath)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return fmt.Errorf("このフォルダは空です")
	}

	// ファイルの読み込み
	var lines []string
	for _, fPath := range files {
		f, err := c.fr.ReadAll(fPath)
		if err != nil {
			return err
		}
		defer f.Close()

		s := bufio.NewScanner(f)
		var lines []string
		for s.Scan() {
			// 置換
			buf := s.Text()
			lines = append(lines, buf)
		}
	}

	// ファイル保存

}
