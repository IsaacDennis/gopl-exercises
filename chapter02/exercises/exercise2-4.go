package exercises

func PopCount(x uint64) int {
	var pop int
	for i := 0; i < 64; i++ {
		pop += int((x >> i)) & 1
	}
	return pop
}
