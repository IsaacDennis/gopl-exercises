package main

import "fmt"

type Point struct {
	X, Y int
}
type Circle struct {
	Point
	Radius int
}

func main() {
	c := Circle{
		Point:  Point{4, 7},
		Radius: 2,
	}
	fmt.Printf("%#v\n", c)
}
