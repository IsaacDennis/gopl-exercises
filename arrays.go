package main

import "fmt"

func main() {
	var b [2]int = [2]int{4, 7}
	for i, num := range b {
		fmt.Printf("%d:%d\n", i, num)
	}
}
