package main

import (
	"fmt"
	"log"
)

type House struct {
	numOfWindows int
	numOfDoors   int
	hasGarage    bool
}

type HouseBuilder struct {
	house House
}

func (b *HouseBuilder) BuildWindows() *HouseBuilder {
	b.house.numOfWindows++
	return b
}

func (b *HouseBuilder) BuildDoors() *HouseBuilder {
	b.house.numOfDoors++
	return b
}

func (b *HouseBuilder) SetGarage(hasGarage bool) *HouseBuilder {
	b.house.hasGarage = hasGarage
	return b
}

func (b *HouseBuilder) Build() (*House, error) {
	if b.house.hasGarage == false {
		return nil, fmt.Errorf("House must have a garage")
	}
	return &b.house, nil
}

func main() {
	hb := &HouseBuilder{}
	house, err := hb.BuildWindows().
		BuildDoors().
		SetGarage(true).
		Build()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(house)
}
