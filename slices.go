package main

import "fmt"

func main() {
	a := [4]int{0, 6, 7, 2}
	b := a[0:3]
	c := a[1:]
	d := a[:3]
	e := a[:]
	fmt.Println(b) // [0 6 7]
	fmt.Println(c) // [6 7 2]
	fmt.Println(d) // [0 6 7]
	fmt.Println(e) // [0 6 7 2]
	f := []int{0, 2, 4, 5}
	fmt.Println(f)
}
