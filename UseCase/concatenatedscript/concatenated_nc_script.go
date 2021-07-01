package concatenatedscript

import "nc-script-converter/Domain/alterationncscript"

type ConcatenatedNcScriptUseCase struct {
	concat alterationncscript.CombinedNcScript
}

func NewConcatenatedNcScriptUseCase(concat alterationncscript.CombinedNcScript) *ConcatenatedNcScriptUseCase {
	return &ConcatenatedNcScriptUseCase{
		concat: concat,
	}
}

func (c *ConcatenatedNcScriptUseCase) ConcatenatedNcScript(inPath string, outPath string) error {
	// そのままドメイン層に委譲
	return c.concat.CombineNcScript(inPath, outPath)
}
