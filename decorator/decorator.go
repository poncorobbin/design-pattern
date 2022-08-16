package main

import "fmt"

type pizza interface {
	getPrice() float64
}

type veggiePizza struct{}

func (v veggiePizza) getPrice() float64 {
	return 10
}

// decorator to change core bevahiour based on topping
type cheeseTopping struct {
	pizza pizza
}

func (c cheeseTopping) getPrice() float64 {
	return c.pizza.getPrice() + 20
}

type tommatoTopping struct {
	pizza pizza
}

func (t tommatoTopping) getPrice() float64 {
	return t.pizza.getPrice() + 5
}

func main() {
	pizza := veggiePizza{}
	fmt.Println("base price:", pizza.getPrice())

	pizzaWithCheese := cheeseTopping{pizza: pizza}
	fmt.Println("adding cheese topping:", pizzaWithCheese.getPrice())

	pizzaWithCheeseAndTomato := tommatoTopping{pizza: pizzaWithCheese}
	fmt.Println("adding tomato and cheese topping:", pizzaWithCheeseAndTomato.getPrice())

	fmt.Println("======")
	pizzaWithTomato := tommatoTopping{pizza: pizza}
	fmt.Println("adding tomato topping:", pizzaWithTomato.getPrice())
}
