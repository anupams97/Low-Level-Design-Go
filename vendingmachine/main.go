package main

import (
	"fmt"
	"github.com/anupams97/Low-Level-Design-Go/vendingmachine/app"
)

func main() {
	vm := app.NewVendingMachine()

	p1 := &app.Product{ProductType: app.Snack, Price: 10}
	p2 := &app.Product{ProductType: app.Drink, Price: 20}
	p3 := &app.Product{ProductType: app.Milk, Price: 30}

	vm.Inventory.AddProduct(p1, 10)
	vm.Inventory.AddProduct(p2, 5)
	vm.Inventory.AddProduct(p3, 1)
	fmt.Println(*vm.Inventory)
	vm.SelectProduct(p1)
	vm.AddCoin(app.TEN)
	vm.Dispense()
	vm.ReturnChange()
	fmt.Println(*vm.Inventory)

	vm.SelectProduct(p1)
	vm.AddCoin(app.FIVE)
	vm.AddCoin(app.TWO)
	vm.AddCoin(app.FIVE)
	vm.Dispense()
	vm.ReturnChange()
	fmt.Println(*vm.Inventory)

	vm.SelectProduct(p3)
	vm.AddCoin(app.TEN)
	vm.AddCoin(app.TEN)
	vm.AddCoin(app.TEN)
	vm.Dispense()
	vm.ReturnChange()
	fmt.Println(*vm.Inventory)

	//item finished
	vm.SelectProduct(p3)
	vm.AddCoin(app.TEN)
	vm.AddCoin(app.TEN)
	vm.AddCoin(app.TEN)
	vm.Dispense()
	vm.ReturnChange()
	fmt.Println(*vm.Inventory)

}
