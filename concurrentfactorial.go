package main
import (
	"strconv"
	"os"
	"fmt"
	"time"
)
func main(){
	start := time.Now()
	channel := make(chan string)
	for _, arg := range os.Args[1:] {
		number, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		go factorial(number, channel)
	}
	for range os.Args[1:] {
		fmt.Println(<-channel)
	}
	fmt.Printf("Tempo para calcular todos os fatoriais: %.2fs", time.Since(start).Seconds())
}
func factorial(num int, ch chan<- string){
	if num < 0 {
		ch <- "Erro: não é possível fatorial de número negativo"
		return
	}
	result := 1
	for i := 1; i <= num; i++ {
		result = result * i
	}
	ch <- fmt.Sprintf("Fatorial de %d: %d", num, result)
}
