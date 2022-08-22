package main

import "fmt"

// element interface
type shape interface {
	getType() string
	accept(visitor)
}

// concrete element
type square struct {
	side int
}

func (s *square) getType() string {
	return "square"
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

// concrete element
type circle struct {
	radius int
}

func (c *circle) getType() string {
	return "circle"
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

// visitor interface
type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
}

// concrete visitor
type areaCalculator struct {
}

func (a *areaCalculator) visitForSquare(*square) {
	fmt.Println("calculating area for square")
}

func (a *areaCalculator) visitForCircle(*circle) {
	fmt.Println("calculating area for circle")
}

// another concrete visitor
type middleCoordinate struct {
}

func (a *middleCoordinate) visitForSquare(*square) {
	fmt.Println("calculating middle point for square")
}

func (a *middleCoordinate) visitForCircle(*circle) {
	fmt.Println("calculating middle point for circle")
}

func main() {
	square := &square{side: 4}
	circle := &circle{radius: 10}

	// add behaviour of shape (element interface) without modify any code
	getArea := &areaCalculator{}

	square.accept(getArea)
	circle.accept(getArea)

	fmt.Println()
	// add another behaviour of shape (element interface)
	middleCoordinate := &middleCoordinate{}

	square.accept(middleCoordinate)
	circle.accept(middleCoordinate)
}
