package exercises

import "io"

type Writer struct {
	count int64
	w     io.Writer
}

func (writer *Writer) Write(p []byte) (int, error) {
	writer.count += int64(len(p))
	return writer.w.Write(p)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var writer Writer
	writer.w = w
	return &writer, &writer.count
}
