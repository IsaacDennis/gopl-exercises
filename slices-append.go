package main

import "fmt"

func main() {
	a := []int{3, 5, 8}
	fmt.Println(a)
	b := append(a, 9)
	fmt.Println(b)
}
