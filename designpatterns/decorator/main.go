package main

import (
	"errors"
	"fmt"
	f "github.com/anupams97/Low-Level-Design-Go/designpatterns/decorator/firstexample"
	s "github.com/anupams97/Low-Level-Design-Go/designpatterns/decorator/secondexample"
	"math/rand/v2"
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

	fn := Decorate2x(getRandom)
	fmt.Println(fn(), fn())

	err := withRetry(doSomething, 10)()
	if err != nil {
		fmt.Println("Final error:", err)
	}

}

func getRandom() int {
	r := rand.IntN(10)
	fmt.Printf("Log r=%d\n", r)
	return r
}

func Decorate2x(fn func() int) func() int {
	return func() int {
		return fn() * 2
	}
}

func doSomething() error {
	r := rand.IntN(10)
	if r > 8 {
		fmt.Println("success")
		return nil
	}
	return errors.New("failed")
}

func withRetry(fn func() error, attempts int) func() error {
	return func() error {
		var lastErr error
		for i := 0; i < attempts; i++ {
			err := fn()
			if err == nil {
				return nil // success
			}
			fmt.Printf("Attempt %d failed: %v\n", i, err)
			lastErr = err
		}
		return fmt.Errorf("all %d attempts failed: %w", attempts, lastErr)
	}
}
