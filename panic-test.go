package main

func main() {
	panicAbove6(7)
}
func panicAbove6(n int) int {
	if n > 6 {
		panic("Número maior que 6")
	}
	return n
}
