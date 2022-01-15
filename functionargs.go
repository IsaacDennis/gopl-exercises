package main

import "fmt"

func main() {
	a := [2]int{4, 7}
	for _, num := range a {
		fmt.Println(num)
	}
	resetArray(&a)
	for _, num := range a {
		fmt.Println(num)
	}
}
func resetArray(arr *[2]int) {
	*arr = [2]int{}
}
