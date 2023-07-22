package exercises

import (
	"bufio"
	"bytes"
)

type LineCounter int
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	var count int
	reader := bytes.NewReader(p)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	*c += WordCounter(count)
	return count, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	var count int
	reader := bytes.NewReader(p)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		count++
	}
	*c += LineCounter(count)
	return count, nil
}
