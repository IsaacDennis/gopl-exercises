package exercises

func rotateLeft(s []int, n int) {
	lastIndex := len(s) - 1
	for i := 0; i < n; i++ {
		first := s[0]
		for j := 1; j <= lastIndex; j++ {
			s[j-1] = s[j]
		}
		s[lastIndex] = first
	}
}
