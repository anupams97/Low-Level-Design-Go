package main

import (
	"fmt"
	app2 "github.com/anupams97/Low-Level-Design-Go/lld/vendingmachine/app"
)

func main() {
	vm := app2.NewVendingMachine()

	p1 := &app2.Product{ProductType: app2.Snack, Price: 10}
	p2 := &app2.Product{ProductType: app2.Drink, Price: 20}
	p3 := &app2.Product{ProductType: app2.Milk, Price: 30}

	vm.Inventory.AddProduct(p1, 10)
	vm.Inventory.AddProduct(p2, 5)
	vm.Inventory.AddProduct(p3, 1)
	fmt.Println(*vm.Inventory)
	vm.SelectProduct(p1)
	vm.AddCoin(app2.TEN)
	vm.Dispense()
	vm.ReturnChange()
	fmt.Println(*vm.Inventory)

	vm.SelectProduct(p1)
	vm.AddCoin(app2.FIVE)
	vm.AddCoin(app2.TWO)
	vm.AddCoin(app2.FIVE)
	vm.Dispense()
	vm.ReturnChange()
	fmt.Println(*vm.Inventory)

	vm.SelectProduct(p3)
	vm.AddCoin(app2.TEN)
	vm.AddCoin(app2.TEN)
	vm.AddCoin(app2.TEN)
	vm.Dispense()
	vm.ReturnChange()
	fmt.Println(*vm.Inventory)

	//item finished
	vm.SelectProduct(p3)
	vm.AddCoin(app2.TEN)
	vm.AddCoin(app2.TEN)
	vm.AddCoin(app2.TEN)
	vm.Dispense()
	vm.ReturnChange()
	fmt.Println(*vm.Inventory)

}
