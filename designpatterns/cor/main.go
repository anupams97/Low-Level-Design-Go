package main

import "fmt"

type handler interface {
	handle(request int)
	setNextHandler(handler handler)
}

type denominationHandler struct {
	denomination int
	nextHandler  handler
}

func (d *denominationHandler) handle(money int) {
	count := money / d.denomination
	fmt.Printf("%d Notes: %d\n", d.denomination, count)
	if d.nextHandler != nil {
		d.nextHandler.handle(money - d.denomination*count)
	}
}

func (d *denominationHandler) setNextHandler(next handler) {
	d.nextHandler = next
}

func main() {
	a := 900
	h500 := &denominationHandler{denomination: 500}
	h100 := &denominationHandler{denomination: 100}

	h500.setNextHandler(h100)

	h500.handle(a)
}
