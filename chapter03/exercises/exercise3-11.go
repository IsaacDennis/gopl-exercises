package main

import (
	"fmt"
	"strings"
)

func isFloat(str string) bool {
	return strings.Contains(str, ".")
}

func hasSign(str string) bool {
	return str[0] == '+' || str[0] == '-'
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	if isFloat(s) {
		dotIndex := strings.Index(s, ".")
		return comma(s[:dotIndex]) + s[dotIndex:]
	}

	i := 0
	sign := ""
	if hasSign(s) {
		i = 1 // Start the first slice with index 1, ignoring the sign byte
		sign = string(s[0])
	}
	return sign + comma(s[i:n-3]) + "," + s[n-3:]
}

func main() {
	n1 := "+123456.789"
	fmt.Printf("%s\n", comma(n1))
}
