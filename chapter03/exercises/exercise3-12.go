package exercises

import (
	"fmt"
)

func main() {
	str1 := "ab"
	str2 := "bb"

	fmt.Printf("%s and %s are anagrams: %v\n", str1, str2, areAnagram(str1, str2))
}

func areAnagram(str1, str2 string) bool {
	str1Map := make(map[rune]int)
	str2Map := make(map[rune]int)

	for _, char := range str1 {
		str1Map[char]++
	}
	for _, char := range str2 {
		str2Map[char]++
	}

	if len(str1Map) != len(str2Map) {
		return false
	}
	for char := range str1Map {
		if str1Map[char] != str2Map[char] {
			return false
		}
	}
	return true
}
