package alterationncscript

import (
	"fmt"
	"log"
	"path/filepath"
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

func (c *CombinedNcScript) CombineNcScript(inPath string, inFiles []string, outPath string) error {
	// ファイルの読み込み
	var conLine []string
	for i, file := range inFiles {
		log.Println("info:ファイル存在確認 file:", inPath, file)
		f := filepath.Join(inPath, file)
		if !(*c.fr).FileExist(f) {
			log.Println("ファイルが存在しません", f)
			return fmt.Errorf("ファイルが存在しません %s", f)
		}

		log.Printf("info: ファイルを読み込みます[%d/%d] file: %s\n", i+1, len(inFiles), file)
		lines, err := (*c.fr).ReadAll(f)
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
	conLine = append(conLine, []string{"M30", "%"}...)

	// 結合ファイル保存
	log.Println("info: 結合ファイルを保存します", conLine)
	if err := (*c.fw).WriteAll(outPath, conLine); err != nil {
		return fmt.Errorf("結合ファイルの保存に失敗しました")
	}

	return nil
}
