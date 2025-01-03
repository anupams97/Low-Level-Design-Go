package lot

import (
	"github.com/anupams97/Low-Level-Design-Go/parkinglot/vehicle"
	"sync"
)

var (
	parkingLotInstance *ParkingLot
	once               sync.Once
)

type ParkingLot struct {
	parkingSystem *ParkingSystem
}

func (l *ParkingLot) ParkVehicle(vehicle vehicle.Vehicle) error {
	err := l.parkingSystem.Park(vehicle)
	if err != nil {
		return err
	}
	return nil
}
func (l *ParkingLot) UnParkVehicle(vehicle vehicle.Vehicle) error {
	err := l.parkingSystem.UnPark(vehicle)
	if err != nil {
		return err
	}
	return nil
}

func GetParkingLot(levelConfig []map[ParkingSpotType]int) *ParkingLot {
	once.Do(func() {
		parkingLotInstance = &ParkingLot{}
		parkingLotInstance.parkingSystem = NewParkingManager(&getFirstFreeSpotStrategy{}, levelConfig)

	})
	return parkingLotInstance
}
