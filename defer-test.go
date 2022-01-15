package main

import "fmt"

func main() {
	defer fmt.Println("Mundo")
	fmt.Print("Olá ")
}
