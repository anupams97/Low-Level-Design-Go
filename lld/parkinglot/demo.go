package main

import (
	"fmt"
	lot2 "github.com/anupams97/Low-Level-Design-Go/lld/parkinglot/lot"
	vehicle2 "github.com/anupams97/Low-Level-Design-Go/lld/parkinglot/vehicle"
)

func main() {
	car := vehicle2.NewCar("UP321234")
	bike1 := vehicle2.NewBike("UP3799434")
	bike2 := vehicle2.NewBike("UP3199434")
	bike3 := vehicle2.NewBike("UP3199431")

	config := []map[lot2.ParkingSpotType]int{
		{
			lot2.COMPACT:    10,
			lot2.TWOWHEELER: 2,
		},
		{
			lot2.COMPACT:    10,
			lot2.TWOWHEELER: 1,
		},
	}

	lot := lot2.GetParkingLot(config)
	err := lot.ParkVehicle(car)
	if err != nil {
		fmt.Println(err)
	}
	err = lot.ParkVehicle(bike1)
	if err != nil {
		fmt.Println(err)
	}
	err = lot.ParkVehicle(bike2)
	if err != nil {
		fmt.Println(err)
	}
	err = lot.ParkVehicle(bike3)
	if err != nil {
		fmt.Println(err)
	}
	err = lot.UnParkVehicle(bike3)
	if err != nil {
		fmt.Println(err)
	}
	bike4 := vehicle2.NewBike("UP31499431")
	err = lot.ParkVehicle(bike4)
	if err != nil {
		fmt.Println(err)
	}

}
