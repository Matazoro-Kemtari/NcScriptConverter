package concatenatedscript

import "nc-script-converter/Domain/alterationncscript"

type ConcatenatedNcScriptUseCase struct {
	concat        *alterationncscript.CombinedNcScript
	readDirectory *alterationncscript.DirViewer
}

func NewConcatenatedNcScriptUseCase(
	concat *alterationncscript.CombinedNcScript,
	readDirectory *alterationncscript.DirViewer,
) *ConcatenatedNcScriptUseCase {
	return &ConcatenatedNcScriptUseCase{
		concat:        concat,
		readDirectory: readDirectory,
	}
}

func (c *ConcatenatedNcScriptUseCase) ConcatenatedNcScript(inPath string, inFiles []string, outPath string, canOpenReview bool) error {
	// そのままドメイン層に委譲
	return c.concat.CombineNcScript(inPath, inFiles, outPath, canOpenReview)
}

func (c *ConcatenatedNcScriptUseCase) DirectoryExist(dirPath string) bool {
	return (*c.readDirectory).DirExist(dirPath)
}

func (c *ConcatenatedNcScriptUseCase) FetchFileNames(dirPath string) ([]string, error) {
	return (*c.readDirectory).FetchDir(dirPath)
}
