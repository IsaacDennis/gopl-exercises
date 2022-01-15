package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	pessoa1 := Person{
		Name: "João",
		Age:  17,
	}
	fmt.Println(pessoa1)
}
