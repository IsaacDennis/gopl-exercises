package exercises

import "fmt"

func comma(str string) string {
	s := str
	for i := len(s); i > 3; i = i - 3 {
		s = s[:i-3] + "," + s[i-3:]
	}
	return s
}

func main() {
	n1 := "12345678"
	fmt.Printf("%s\n", comma(n1))
}
