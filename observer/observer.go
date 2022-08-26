package main

import "fmt"

// subject interface
type Subject interface {
	subscribe(o observer)
	unsubscribe(o observer)
	notifyAll()
}

// concrete subject
type Item struct {
	observerList []observer
	name         string
	inStock      bool
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock \n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *Item) subscribe(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) unsubscribe(o observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *Item) notifyAll() {
	for _, o := range i.observerList {
		o.update(i.name)
	}
}

func removeFromSlice(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

// observer interface
type observer interface {
	update(string)
	getID() string
}

// concrete observer
type customer struct {
	id string
}

func (c customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c customer) getID() string {
	return c.id
}

func main() {
	c1 := customer{id: "ponco@loc.com"}
	c2 := customer{id: "robbi@loc.com"}

	item := newItem("Lenovo")
	item.subscribe(c1)
	item.subscribe(c2)

	item.updateAvailability()
}
