package main

import (
	"fmt"
	"github.com/anupams97/Low-Level-Design-Go/designpatterns/builder/carbuilder/car"
	"github.com/anupams97/Low-Level-Design-Go/designpatterns/builder/studentbuilder/student"
)

func main() {
	dir := student.NewDirector()
	student := dir.Build()
	fmt.Println(*student)

	scBuilder := car.NewSportsCarBuider()
	ocBuilder := car.NewOffroadCarBuider()
	carDir := car.NewDirector(scBuilder)
	fmt.Println(carDir.CreateSportsCar("ferrari", 800))
	carDir = car.NewDirector(ocBuilder)
	fmt.Println(carDir.CreateOffRoadCar("hilux", 1000))
}
