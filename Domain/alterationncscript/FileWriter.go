package alterationncscript

type FileWriter interface {
	WriteAll(path string, contents []string) error
}
