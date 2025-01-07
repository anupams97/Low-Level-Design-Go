package concurrentmarket

import "fmt"

type Traders struct {
	ID   int
	Name string
}

func (t *Traders) OnNotify(event Event) {
	fmt.Printf("Stock: %s | Price: %d | Received by Trader: %s\n",
		event.StockName, event.Price, t.Name)
}
