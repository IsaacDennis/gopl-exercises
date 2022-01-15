package words

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"kayak", true},
		{"ab", false},
	}
	for _, test := range tests {
		if result := IsUniPalindrome(test.input); result != test.expected {
			t.Errorf("IsUniPalindrome(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}
func TestIsUniPalindrome(t *testing.T) {
	if !IsUniPalindrome("été") {
		t.Error(`IsUniPalindrome("été") = false, expected true`)
	}
}
