package main

import "fmt"

// component interface
type Participant interface {
	getName() string
	answer() string
}

// concrete component
type Student struct {
	moderator Moderator
	name      string
}

func (s Student) getName() string {
	return s.name
}

func (s Student) answer() string {
	return fmt.Sprintf("%s menjawab pertanyaan", s.name)
}

func (s Student) pressButton() {
	s.moderator.notify(s)
}

// moderator interface
type Moderator interface {
	notify(sender Participant)
}

// concrete moderator
type Teacher struct {
	P1, P2 Participant
}

func (t Teacher) notify(p Participant) {
	var answer string
	if t.P1.getName() == p.getName() {
		answer = t.P1.answer()
	} else {
		answer = t.P2.answer()
	}
	fmt.Println(answer)
}

func main() {
	t := &Teacher{}

	p1 := Student{
		moderator: t,
		name:      "ponco",
	}

	p2 := Student{
		moderator: t,
		name:      "robbi",
	}

	t.P1 = p1
	t.P2 = p2

	p1.pressButton()
}
