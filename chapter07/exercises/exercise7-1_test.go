package exercises

import "testing"

func TestWordCounter(t *testing.T) {
	var counter WordCounter
	test := []byte("Hello, world! Hello, world!")
	result, _ := counter.Write(test)
	t.Log(result)
}

func TestLineCounter(t *testing.T) {
	var counter LineCounter
	test := []byte("Hello, world!\nHello, world!")
	result, _ := counter.Write(test)
	t.Log(result)
}
