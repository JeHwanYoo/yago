package mock

import (
	"os"
)

type Writer struct {
	data          []byte
	WriteFileFunc func(name string, data []byte, perm os.FileMode) error
}

func (w *Writer) WriteFile(name string, data []byte, perm os.FileMode) error {
	if w.WriteFileFunc != nil {
		return w.WriteFileFunc(name, data, perm)
	}

	w.data = data
	return nil
}

func (w *Writer) GetWrittenData() []byte {
	return w.data
}
