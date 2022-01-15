package main

import (
	"fmt"
	"os"
	"strings"
	
)
func main(){
	var s, sep string
	for _, arg := range os.Args[1:]{
		s += sep + arg
		sep = " "

	}
	fmt.Println(s)
	main2()
	fmt.Println(os.Args[1:])
	main3()
}
// Alternativa à função main utilizando o pacote "strings"
func main2(){
	fmt.Println(strings.Join(os.Args, " "));
}
func main3(){
	for i, arg := range os.Args[1:]{
		fmt.Println(i, arg);
	}
}

