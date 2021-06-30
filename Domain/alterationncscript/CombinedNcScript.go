package alterationncscript

import (
	"fmt"
	"log"
	"sort"
)

type CombinedNcScript struct {
	dir DirViewer
	fr  FileReader
	cnv ConvertedNcScript
}

func NewCombinedNcScript(dir DirViewer, fr FileReader, cnv ConvertedNcScript) *CombinedNcScript {
	return &CombinedNcScript{
		dir: dir,
		fr:  fr,
		cnv: cnv,
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

	// ファイル名順にする
	sort.Slice(files, func(i, j int) bool {
		return files[i] < files[j]
	})

	// ファイルの読み込み
	var conLine []string
	for _, fPath := range files {
		lines, err := c.fr.ReadAll(fPath)
		if err != nil {
			return err
		}

		resLine, err := c.cnv.Convert(lines)
		if err != nil {
			return err
		}
		conLine = append(conLine, resLine...)
	}

	// 前後に"%"追加
	conLine = append([]string{"%"}, conLine...)
	conLine = append(conLine, "%")

	// ファイル保存
	log.Println("info:", conLine)

	return nil
}
