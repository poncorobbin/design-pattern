package main

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct {
	Color Color // this is what bridge pattern is
}

func (c Circle) Draw() {
	if c.Color != nil {
		fmt.Println("drawing circle with color", c.Color.Hex())
	} else {
		fmt.Println("drawing circle")
	}
}

func (c Circle) Clone() Circle {
	return Circle{
		Color: c.Color,
	}
}

type Square struct {
	Color Color
}

func (s Square) Draw() {
	fmt.Println("drawing square")
}

type Color interface {
	Hex() string
}

type Red struct{}

func (r Red) Hex() string {
	return "#a83232"
}

type Blue struct{}

func (b Blue) Hex() string {
	return "#2522d4"
}

func main() {
	c := &Circle{}
	c.Color = Red{}
	c.Draw()

	c2 := c.Clone()
	c2.Color = nil
	c2.Draw()
}
