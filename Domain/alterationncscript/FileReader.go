package alterationncscript

import "os"

type FileReader interface {
	ReadAll(path string) (*os.File, error)
}
