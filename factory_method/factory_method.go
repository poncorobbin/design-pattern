package main

import "fmt"

// product
type Product interface {
	Deliver()
}

type ProductA struct{}

func (p *ProductA) Deliver() {
	fmt.Println("this is deliver from ProductA")
}

type ProductB struct{}

func (p *ProductB) Deliver() {
	fmt.Println("this is deliver from ProductB")
}

type ProductC struct{}

func (p *ProductC) Deliver() {
	fmt.Println("this is deliver from ProductC")
}

// creator
type ProductCreator interface {
	CreateProduct(productType int) Product
}

type CreatorOne struct{}

func (c *CreatorOne) CreateProduct(productType int) Product {
	if productType == 1 {
		return &ProductA{}
	} else if productType == 2 {
		return &ProductC{}
	}
	return nil
}

type CreatorTwo struct{}

func (c *CreatorTwo) CreateProduct(productType int) Product {
	return &ProductB{}
}

func main() {
	var productCreator ProductCreator
	productCreator = &CreatorOne{}

	product := productCreator.CreateProduct(1)
	product.Deliver()
}
