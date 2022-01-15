package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	n, err := failAbove5(6)
	if err != nil {
		log.Fatalf("Erro: %v\n", err)
		return
	}
	fmt.Printf("Número %d\n", n)
}

/* Se o número "n" for maior que 5, um erro ocorre.
Senão, o próprio n é retornado */
func failAbove5(n int) (int, error) {
	if n > 5 {
		return 0, errors.New("Um número maior que 5 foi passado como parâmetro.")
	}
	return n, nil
}
