package exercises

import "testing"

func TestFindMatches(t *testing.T) {
	testString := "$foo$baa $caa $foo"
	matches := make(map[string]bool)
	FindMatches(testString, matches)
	t.Log(matches)
}

func TestExpand(t *testing.T) {
	testString := "$foo$baa $caa $foo"
	expected := "11 1 1"
	fn := func(a string) string {
		return "1"
	}
	expand(&testString, fn)
	if testString != expected {
		t.Errorf("Error: expected %s, got %s", expected, testString)
	}
	t.Log(testString)
}
