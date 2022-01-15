package main

import "fmt"

func main() {
	idades := map[string]int{
		"João":   17,
		"Carlos": 20,
	}
	fmt.Println(idades)
	delete(idades, "João")
	fmt.Println(idades)
	if _, ok := idades["João"]; ok {
		fmt.Println("Existe!")
	} else {
		fmt.Println("Não existe!")
	}
}
