package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }
type Person struct{ age int }
type FlyingCharacter interface {
	Character
	Fly()
	Land()
}
type Character interface {
	Walk()
	Attack()
}

func main() {
	a := Point{6, 8}
	b := Point{3, 4}
	distance := a.Distance(&b)
	fmt.Println(distance)
}
func (p *Point) Distance(q *Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func (p *Person) Age() int {
	return p.age
}
func LandAll(fcs []FlyingCharacter) {
	for _, character := range fcs {
		character.Land()
	}
}
