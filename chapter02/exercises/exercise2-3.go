package exercises

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Devolve a população (número de bits definidos) de x
func PopCount(x uint64) int {
	var pop int
	for i := 0; i < 8; i++ {
		pop += int(pc[byte(x>>(i*8))])
	}
	return pop
}
