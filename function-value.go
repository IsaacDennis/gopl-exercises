package main

import "fmt"

func main() {
	fmt.Println(sumAll(1, 4, 5))
	fmt.Println(execute(sumAll, 1, 4, 5))
}
func execute(f func(...int) int, params ...int) int {
	return f(params...)
}
func sumAll(params ...int) int {
	var sum int
	for _, value := range params {
		sum += value
	}
	return sum
}
