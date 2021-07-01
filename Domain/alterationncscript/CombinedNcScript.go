package alterationncscript

import (
	"fmt"
	"log"
	"sort"
)

type CombinedNcScript struct {
	dir *DirViewer
	fr  *FileReader
	cnv *ConvertedNcScript
	fw  *FileWriter
}

func NewCombinedNcScript(dir *DirViewer, fr *FileReader, cnv *ConvertedNcScript, fw *FileWriter) *CombinedNcScript {
	return &CombinedNcScript{
		dir: dir,
		fr:  fr,
		cnv: cnv,
		fw:  fw,
	}
}

func (c *CombinedNcScript) CombineNcScript(inPath string, outPath string) error {
	// ファイル一覧取得
	log.Printf("info: ファイルの一覧を取得します Folder: %s\n", inPath)
	files, err := (*c.dir).FetchDir(inPath)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return fmt.Errorf("このフォルダは空です")
	}

	// ファイル名順にする
	log.Print("info: 取得したファイル名を並び替えます before: ", files, ",")
	sort.Slice(files, func(i, j int) bool {
		return files[i] < files[j]
	})
	log.Println("after: ", files)

	// ファイルの読み込み
	var conLine []string
	for i, fPath := range files {
		log.Printf("info: ファイルを読み込みます[%d/%d] file: %s\n", i, len(fPath), fPath)
		lines, err := (*c.fr).ReadAll(fPath)
		if err != nil {
			return err
		}

		resLine, err := (*c.cnv).Convert(lines)
		if err != nil {
			return err
		}
		conLine = append(conLine, resLine...)
	}

	// 前後に"%"追加
	conLine = append([]string{"%"}, conLine...)
	conLine = append(conLine, "%")

	// 結合ファイル保存
	log.Println("info: 結合ファイルを保存します", conLine)
	if err := (*c.fw).WriteAll(outPath, conLine); err != nil {
		return fmt.Errorf("結合ファイルの保存に失敗しました")
	}

	return nil
}
