package exercises

func PopCount(x uint64) int {
	word := x
	var pop int
	for word != 0 {
		word = word & (word - 1)
		pop++
	}
	return pop
}
