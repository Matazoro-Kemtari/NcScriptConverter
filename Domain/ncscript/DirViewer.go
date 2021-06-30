package ncscript

type DirViewer interface {
	FetchDir(path string) ([]string, error)
}
