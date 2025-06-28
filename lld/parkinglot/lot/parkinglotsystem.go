package lot

import (
	"errors"
	"fmt"
	"github.com/anupams97/Low-Level-Design-Go/lld/parkinglot/vehicle"
)

type ParkingSystem struct {
	levels   []*ParkingLevel
	strategy ParkingStrategy
	occupied map[string]*ParkingSpot
}

func NewParkingManager(strategy ParkingStrategy, levelConfig []map[ParkingSpotType]int) *ParkingSystem {
	ps := &ParkingSystem{strategy: strategy}
	for i, lc := range levelConfig {
		ps.levels = append(ps.levels, NewParkingLevel(i, lc))
	}
	ps.occupied = make(map[string]*ParkingSpot)
	return ps
}
func (p *ParkingSystem) AllocateSpot(vehicle vehicle.Vehicle) (*ParkingSpot, error) {
	spot := p.strategy.AllocateSpot(p.levels, vehicle.GetVehicleType())
	if spot != nil {
		return spot, nil
	}
	return nil, errors.New("cannot find parking spot")

}
func (p *ParkingSystem) setStrategy() {
	p.strategy = &getFirstFreeSpotStrategy{}
}

type getFirstFreeSpotStrategy struct {
}

func (g *getFirstFreeSpotStrategy) AllocateSpot(levels []*ParkingLevel, vehicleType vehicle.VehicleType) *ParkingSpot {
	for _, level := range levels {
		for _, spot := range level.parkingSpots {
			if spot.IsAvailable() && spot.isCompatible(vehicleType) {
				return spot
			}
		}
	}
	return nil
}

func (p *ParkingSystem) Park(vehicle vehicle.Vehicle) error {
	spot := p.strategy.AllocateSpot(p.levels, vehicle.GetVehicleType())
	if spot != nil {
		err := spot.Park(vehicle)
		if err != nil {
			return fmt.Errorf("cannot park err: %v", err)
		}
		p.occupied[vehicle.GetLicensePlate()] = spot
		fmt.Printf("vehicle %v parked at spot %d level %d \n", vehicle.GetLicensePlate(), spot.GetSpotID(), spot.GetLevel())
		return nil
	}
	return errors.New("cannot find spot")
}

func (p *ParkingSystem) UnPark(vehicle vehicle.Vehicle) error {
	if spot, ok := p.occupied[vehicle.GetLicensePlate()]; ok {
		err := spot.UnPark()
		if err != nil {
			return fmt.Errorf("cannot unpark err: %v", err)
		}
		fmt.Printf("vehicle %v unparked from spot %d level %d \n", vehicle.GetLicensePlate(), spot.GetSpotID(), spot.GetLevel())
		return nil
	}
	return errors.New("cannot find vehicle in parking spot")
}
