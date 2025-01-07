package main

import (
	"fmt"
	"github.com/anupams97/Low-Level-Design-Go/designpatterns/observer/concurrent/concurrentmarket"
)

func main() {
	sm := concurrentmarket.NewStockMarket()
	t1 := concurrentmarket.Traders{ID: 1, Name: "anupam"}
	t2 := concurrentmarket.Traders{ID: 2, Name: "someone"}

	err := sm.Register(&t1)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = sm.Register(&t2)
	if err != nil {
		fmt.Println(err)
		return
	}

	sm.Notify(concurrentmarket.Event{StockName: "google", Price: 1000})
	sm.Notify(concurrentmarket.Event{StockName: "google", Price: 1200})
	sm.DeRegister(&t2)
	sm.Notify(concurrentmarket.Event{StockName: "grab", Price: 100})
	sm.Notify(concurrentmarket.Event{StockName: "grab", Price: 200})
	sm.Wg.Wait()
}
