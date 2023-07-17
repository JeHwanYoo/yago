package mock

import (
	"os"
)

type Writer struct {
	data []byte
}

func (w *Writer) WriteFile(_ string, data []byte, _ os.FileMode) error {
	w.data = data
	return nil
}

func (w *Writer) GetWrittenData() []byte {
	return w.data
}
