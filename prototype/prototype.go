package main

import "fmt"

type Prototype interface {
	Clone() Prototype
}

type Robot struct {
	ID       int
	Name     string
	numOfCpu int
}

func (r *Robot) Clone() Prototype {
	return &Robot{
		ID:       r.ID,
		Name:     r.Name,
		numOfCpu: r.numOfCpu,
	}
}

func main() {
	r1 := &Robot{1, "ponco", 2}
	r2, ok := r1.Clone().(*Robot)
	if ok {
		r2.Name = "robbi"
		r2.numOfCpu = 5
	}

	fmt.Printf("r1: %v\n", r1)
	fmt.Printf("r2: %v\n", r2)
}
