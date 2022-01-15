package main

import "fmt"

func main() {
	r := [...]int{99: 2}
	for _, num := range r {
		fmt.Println(num)
	}
}
