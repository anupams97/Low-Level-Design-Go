package main

import (
	"github.com/anupams97/Low-Level-Design-Go/designpatterns/observer/simpleobserver/stockmarket"
)

func main() {
	stockMarket := &stockmarket.StockMarket{Traders: make(map[stockmarket.Observer]struct{})}
	trader1 := &stockmarket.Traders{Name: "anupam", ID: 1}
	trader2 := &stockmarket.Traders{Name: "someone", ID: 2}

	stockMarket.Register(trader1)
	stockMarket.Register(trader2)
	stockMarket.Notify(stockmarket.Event{EventID: 1, Info: "some info 1"})
	stockMarket.Notify(stockmarket.Event{EventID: 2, Info: "some info 2"})

	stockMarket.DeRegister(trader2)
	stockMarket.Notify(stockmarket.Event{EventID: 3, Info: "some info 3"})

}
