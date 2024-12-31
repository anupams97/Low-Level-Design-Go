package main

type Memento struct {
	state string
}

func (m *Memento) getState() string {
	return m.state
}
