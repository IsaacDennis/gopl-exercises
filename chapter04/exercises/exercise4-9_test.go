package exercises

import (
	"os"
	"testing"
)

func TestWordFreq(t *testing.T) {
	file, _ := os.Open("testfile.txt")
	freq := wordfreq(file)
	t.Log(freq)
}
