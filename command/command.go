package main

import "fmt"

// command interface
type command interface {
	execute()
}

// on command
type onCommand struct {
	device device
}

func (on onCommand) execute() {
	on.device.on()
}

// off command
type offCommand struct {
	device device
}

func (off offCommand) execute() {
	off.device.off()
}

// invoker
type button struct {
	command command
}

func (b button) press() {
	b.command.execute()
}

// receiver interface
type device interface {
	on()
	off()
}

type tv struct{}

func (t tv) on() {
	fmt.Println("turning on tv")
}

func (t tv) off() {
	fmt.Println("turning off tv")
}

func main() {
	tv := &tv{}

	on := onCommand{device: tv}
	btn := button{command: &on}
	btn.press()

	off := offCommand{device: tv}
	btn = button{command: &off}
	btn.press()
}
