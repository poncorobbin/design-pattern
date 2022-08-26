package main

import "fmt"

// iterator interface
type Iterator interface {
	hasNext() bool
	getNext() *User // it can be generic
}

// concrete iterator
type UserIterator struct {
	idx   int
	users []*User
}

func (u *UserIterator) hasNext() bool {
	if u.idx < len(u.users) {
		return true
	}
	return false
}

func (u *UserIterator) getNext() *User {
	if u.hasNext() {
		user := u.users[u.idx]
		u.idx++
		return user
	}
	return nil
}

// collection interface
type Collection interface {
	createIterator() Iterator
}

// concrete collection
type UserCollection struct {
	users []*User
}

func (u *UserCollection) createIterator() Iterator {
	return &UserIterator{users: u.users}
}

type User struct {
	name string
	age  int
}

func main() {
	u1 := &User{name: "ponco", age: 25}
	u2 := &User{name: "robbi", age: 30}

	userCollection := UserCollection{users: []*User{u1, u2}}
	it := userCollection.createIterator()

	for it.hasNext() {
		user := it.getNext()
		fmt.Println(user.name)
		fmt.Println(user.age)
	}
}
