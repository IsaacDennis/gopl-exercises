package exercises

import "strings"

// Using map as set
func FindMatches(s string, matches map[string]bool) {
	for i, char := range s {
		if char == '$' {
			for j := i + 1; j < len(s); j++ {
				if s[j] == ' ' || s[j] == '$' {
					matches[s[i:j]] = true
					break
				} else if j == len(s)-1 {
					matches[s[i:]] = true
				}
			}
		}
	}
}

func expand(s *string, f func(string) string) {
	matches := make(map[string]bool)
	FindMatches(*s, matches)
	for match := range matches {
		*s = strings.ReplaceAll(*s, match, f(match[1:]))
	}
}
