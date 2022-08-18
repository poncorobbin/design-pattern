package main

import (
	"fmt"
	"strings"
)

type handler interface {
	handle(string)
	setNext(handler)
}

// lower case
type lowerCaseHandler struct {
	next handler
}

func (lc lowerCaseHandler) handle(s string) {
	s = strings.ToLower(s)
	fmt.Println("lower", s)

	if lc.next != nil {
		lc.next.handle(s)
	}
}

func (lc *lowerCaseHandler) setNext(h handler) {
	lc.next = h
}

// remove white space
type removeWhiteSpaceHandler struct {
	next handler
}

func (r removeWhiteSpaceHandler) handle(s string) {
	s = strings.Replace(s, " ", "", -1)
	fmt.Println("remove white", s)

	if r.next != nil {
		r.next.handle(s)
	}
}

func (r *removeWhiteSpaceHandler) setNext(h handler) {
	r.next = h
}

func main() {
	str := "Ponco Robbi"
	lc := &lowerCaseHandler{}
	rw := &removeWhiteSpaceHandler{}

	lc.setNext(rw)
	lc.handle(str)
}
