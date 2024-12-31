package main

import (
	"awesomeProject/VisitorPattern/shape"
	"fmt"
)

func main() {
	s := shape.Square{}
	s.SetSide(10)

	c := shape.Circle{}
	c.SetRadius(10)

	areaCal := &shape.Area{}
	s.Accept(areaCal)
	fmt.Printf("Area of Square: %d \n", areaCal.Area)
	c.Accept(areaCal)
	fmt.Printf("Area of Circle: %d \n", areaCal.Area)

	periCal := &shape.Perimeter{}
	s.Accept(periCal)
	fmt.Printf("Peri of Square: %d \n", periCal.Perimeter)
	c.Accept(periCal)
	fmt.Printf("Peri of Circle: %d \n", periCal.Perimeter)

}
