package main

type Originator struct {
	state string
}

func (o *Originator) createMemento() *Memento {
	return &Memento{state: o.state}
}

func (o *Originator) restore(memento *Memento) {
	o.state = memento.state
}

func (o *Originator) setState(s string) {
	o.state = s
}
