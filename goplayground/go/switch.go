package _go

import "fmt"

func mainN() {
	whoami := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int, float64:
			fmt.Println("I'm a number")
		default:
			fmt.Println("I don't know")
		}
	}
	whoami(1)
	whoami(1.0)
	whoami(true)
	whoami("hello")
}
