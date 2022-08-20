package main

import "fmt"

type state interface {
	save()
}

type Document struct {
	state state
}

func newDocument(isAdmin bool) *Document {
	d := &Document{}
    if isAdmin {
	    d.changeState(&Publish{d})
    } else {
	    d.changeState(&Draft{d})
    }
	return d
}

func (d *Document) save() {
	d.state.save()
}

func (d *Document) changeState(state state) {
	d.state = state
}

// draft state
type Draft struct {
	document *Document
}

func (dr Draft) save() {
	fmt.Println("current state is draft, so publish it")
	dr.document.changeState(&Publish{document: dr.document})
}

// publish state
type Publish struct {
	document *Document
}

func (p Publish) save() {
	fmt.Println("current state is published, so save it as draft")
	p.document.changeState(&Draft{document: p.document})
}

func main() {
	doc := newDocument(true)
	doc.save()
}
