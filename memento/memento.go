package main

import "fmt"

// memento
type memento struct {
	text string
}

func (m memento) getText() string {
	return m.text
}

// originator
type editor struct {
	text string
}

func (e *editor) save() memento {
	return memento{text: e.text}
}

func (e *editor) restore(m memento) {
	e.text = m.getText()
}

// caretaker
type controller struct {
	editor    *editor
	histories []memento
	ptr       int // current index of snapshot
}

func (c *controller) cmdSave() {
	m := c.editor.save()
	c.histories = append(c.histories, m)
	c.ptr = len(c.histories) - 1
}

func (c *controller) cmdUndo() {
	if c.ptr <= 0 {
		return
	}
	c.ptr -= 1
	m := c.histories[c.ptr]
	c.editor.restore(m)
}

func (c *controller) cmdRedo() {
	l := len(c.histories) - 1
	if c.ptr >= l {
		return
	}

	c.ptr++
	m := c.histories[c.ptr]
	c.editor.restore(m)
}

func main() {
	editor := editor{}

	c := controller{}
	c.editor = &editor

	editor.text = "ponco"
	c.cmdSave() // snapshot 1
	fmt.Println(editor.text)

	editor.text = "robbi"
	c.cmdSave() // snapshot 2
	fmt.Println(editor.text)

	c.cmdUndo() // snapshot 1
	fmt.Println(editor.text)

	c.cmdRedo() // snapshot 2
	fmt.Println(editor.text)

	c.cmdRedo() // snapshot 2
	fmt.Println(editor.text)
}
