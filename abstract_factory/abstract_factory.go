package main

import "fmt"

type Price interface {
	Price() float64
}

// products interface
type Chair interface {
	Price
	HasLegs() bool
}

type Sofa interface {
	Price
	IsSoft() bool
}

type FurnitureFactory interface { //abstract factory
	CreateChair() Chair
	CreateSofa() Sofa
}

// victorian product
type VictorianChair struct{}

func (v *VictorianChair) Price() float64 {
	return 1000000
}

func (v *VictorianChair) HasLegs() bool {
	return true
}

// victorian product
type VictorianSofa struct{}

func (v *VictorianSofa) Price() float64 {
	return 1000000
}

func (v *VictorianSofa) IsSoft() bool {
	return true
}

// victorian factory
type VictorianFactory struct{}

func (v VictorianFactory) CreateChair() Chair {
	return &VictorianChair{}
}

func (v VictorianFactory) CreateSofa() Sofa {
	return &VictorianSofa{}
}

// modern product
type ModernChair struct{}

func (m *ModernChair) Price() float64 {
	return 2000000
}

func (m *ModernChair) HasLegs() bool {
	return false
}

// modern factory
type ModernFactory struct{}

func (m *ModernFactory) CreateChair() Chair {
	return &ModernChair{}
}

func (m *ModernFactory) CreateSofa() Sofa {
	return nil
}

func main() {
	var furnitureFactory FurnitureFactory
	furnitureFactory = &VictorianFactory{}

	chair := furnitureFactory.CreateChair()
	fmt.Println("VictorianChair", chair.Price())
	sofa := furnitureFactory.CreateSofa()
	fmt.Println("VictorianSofa", sofa.Price())

	furnitureFactory = &ModernFactory{}
	chair = furnitureFactory.CreateChair()
	fmt.Println("modernChair", chair.Price())
}
