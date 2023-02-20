package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
)

func main() {
	symbol := [...]string{USD: "$", EUR: "€"}
	fmt.Println(symbol[EUR]) // €
}
