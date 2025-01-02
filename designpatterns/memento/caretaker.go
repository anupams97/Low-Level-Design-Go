package main

type Caretaker struct {
	history []*Memento
}

func (c *Caretaker) addMemento(m *Memento) {
	c.history = append(c.history, m)

}
func (c *Caretaker) getLastMemento() *Memento {
	return c.history[len(c.history)-1]
}
