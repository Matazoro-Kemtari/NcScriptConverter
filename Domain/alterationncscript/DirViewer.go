package alterationncscript

type DirViewer interface {
	FetchDir(path string) ([]string, error)
	DirExist(path string) bool
}
