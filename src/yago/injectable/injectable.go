package injectable

import "os"

type Parser interface {
	Parse(*[]byte) (*interface{}, error)
}

type Generator interface {
	Generate(*interface{}) (string, error)
}

type Writer interface {
	WriteFile(name string, data *[]byte, perm os.FileMode) error
}
