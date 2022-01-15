package main

import "fmt"

func main() {
	a := []int{3, 7, 9}
	fmt.Println(a)
	insert(87, &a)
	fmt.Println(a)
}
func insert(num int, s *[]int) {
	*s = append(*s, num)
}
