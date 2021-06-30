package alterationncscript

type FileReader interface {
	ReadAll(path string) ([]string, error)
}
