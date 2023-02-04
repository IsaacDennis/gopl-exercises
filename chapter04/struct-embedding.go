package main

import "fmt"

type Point struct {
	X, Y int
}
type Circle struct {
	Center Point
	Radius int
}

func main() {
	circle1 := Circle{
		Center: Point{X: 1, Y: 2},
		Radius: 2,
	}
	fmt.Println(circle1)
}
