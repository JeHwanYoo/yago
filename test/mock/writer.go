package mock

import (
	"os"
)

type Writer struct {
	data []byte
	Err  error
}

func (w *Writer) WriteFile(_ string, data []byte, _ os.FileMode) error {
	w.data = data
	return w.Err
}

func (w *Writer) GetWrittenData() []byte {
	return w.data
}
