package exercises

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	var buffer bytes.Buffer
	writer, count := CountingWriter(&buffer)
	writer.Write([]byte("Hello, world!"))
	t.Log(buffer.String())
	t.Log(count)
}
