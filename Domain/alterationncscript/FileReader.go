package alterationncscript

type FileReader interface {
	ReadAll(path string) ([]string, error)
	FileExist(file string) bool
}
