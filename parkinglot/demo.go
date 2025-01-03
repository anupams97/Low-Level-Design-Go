package main

import (
	"fmt"
	lot "github.com/anupams97/Low-Level-Design-Go/parkinglot/lot"
	"github.com/anupams97/Low-Level-Design-Go/parkinglot/vehicle"
)

func main() {
	car := vehicle.NewCar("UP321234")
	bike1 := vehicle.NewBike("UP3799434")
	bike2 := vehicle.NewBike("UP3199434")
	bike3 := vehicle.NewBike("UP3199431")

	config := []map[lot.ParkingSpotType]int{
		{
			lot.COMPACT:    10,
			lot.TWOWHEELER: 2,
		},
		{
			lot.COMPACT:    10,
			lot.TWOWHEELER: 1,
		},
	}

	lot := lot.GetParkingLot(config)
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
	bike4 := vehicle.NewBike("UP31499431")
	err = lot.ParkVehicle(bike4)
	if err != nil {
		fmt.Println(err)
	}

}
