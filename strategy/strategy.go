package main

import "fmt"

type strategy interface {
	execute(num1, num2 int) int
}

type operationAdd struct{}

func (a *operationAdd) execute(num1, num2 int) int {
	return num1 + num2
}

type operationSubstract struct{}

func (s *operationSubstract) execute(num1, num2 int) int {
	return num1 - num2
}

type operationMultiply struct{}

func (s *operationMultiply) execute(num1, num2 int) int {
	return num1 * num2
}

// this is context
// note : context does not need implement interface
type context struct {
	strategy strategy
}

func newContext(strategy strategy) *context {
	return &context{strategy: strategy}
}

func (c *context) calculate(num1, num2 int) int {
	return c.strategy.execute(num1, num2)
}

func (c *context) setStrategy(strategy strategy) {
	c.strategy = strategy
}

func main() {
	add := operationAdd{}
	ctx := newContext(&add)
	total := ctx.calculate(5, 2)
	fmt.Println(total)

	subs := operationSubstract{}
	ctx.setStrategy(&subs)
	total = ctx.calculate(total, 3)
	fmt.Println(total)

	mult := operationMultiply{}
	ctx.setStrategy(&mult)
	total = ctx.calculate(total, 5)
	fmt.Println(total)
}
