package main

import (
	"fmt"
	f "github.com/anupams97/Low-Level-Design-Go/designpatterns/decorator/firstexample"
	s "github.com/anupams97/Low-Level-Design-Go/designpatterns/decorator/secondexample"
)

func main() {
	pizza := &f.VegMania{}
	fmt.Println("Base Pizza:", pizza, "Price:", pizza.GetPrice())

	cheesePizza := f.NewCheeseTopping(pizza)
	fmt.Println("With Cheese:", cheesePizza, "Price:", cheesePizza.GetPrice())

	sauceCheesePizza := f.NewSauce(cheesePizza)
	fmt.Println("With Sauce & Cheese:", sauceCheesePizza, "Price:", sauceCheesePizza.GetPrice())

	coffee := s.GetCoffee
	fmt.Println("Base Coffee:", coffee()) // normal coffee

	// Apply decorators dynamically
	coffeeWithIceCream := s.AddOns(coffee, "ice cream")
	fmt.Println("With Ice Cream:", coffeeWithIceCream()) // normal coffee + ice cream

	coffeeWithWhippedCream := s.AddOns(coffeeWithIceCream, "whipped cream")
	fmt.Println("With Ice Cream & Whipped Cream:", coffeeWithWhippedCream())

}
