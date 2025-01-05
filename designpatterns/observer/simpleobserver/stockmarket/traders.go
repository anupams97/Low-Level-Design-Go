package stockmarket

import (
	"fmt"
)

type Traders struct {
	ID   int
	Name string
}

func (t *Traders) OnNotify(event Event) {
	fmt.Printf("stock event id: %d info :%s received by trader : %s\n", event.EventID, event.Info, t.Name)
}
