package lot

import (
	"fmt"
	"github.com/anupams97/Low-Level-Design-Go/lld/parkinglot/vehicle"
)

type ParkingState int

const (
	AVAILABLE ParkingState = iota
	OCCUPIED
	UNAVAILABLE
)

type ParkingSpotType int

const (
	COMPACT ParkingSpotType = iota
	LARGE
	TWOWHEELER
	HANDICAP
)

type ParkingSpot struct {
	spotID          int
	level           int
	state           ParkingState
	parkingSpotType ParkingSpotType
	parkedVehicle   vehicle.Vehicle
}

func NewParkingSpot(spotID int, level int, parkingSpotType ParkingSpotType) *ParkingSpot {
	return &ParkingSpot{spotID: spotID, level: level, parkingSpotType: parkingSpotType, state: AVAILABLE}
}
func (p *ParkingSpot) IsAvailable() bool {
	return p.state == AVAILABLE
}

func (p *ParkingSpot) GetSpotID() int {
	return p.spotID
}
func (p *ParkingSpot) GetLevel() int {
	return p.level
}

func (p *ParkingSpot) GetSpotType() ParkingSpotType {
	return p.parkingSpotType
}

func (p *ParkingSpot) Park(vehicle vehicle.Vehicle) error {
	if !p.isCompatible(vehicle.GetVehicleType()) {
		return fmt.Errorf("vehicle type: %v not compatible with parking spot type: %v", vehicle.GetVehicleType(), p.GetSpotType())
	}
	p.parkedVehicle = vehicle
	p.state = OCCUPIED
	return nil
}

func (p *ParkingSpot) UnPark() error {
	if p.parkedVehicle == nil {
		return fmt.Errorf("spot %d level %d is already vacant", p.spotID, p.level)
	}
	p.parkedVehicle = nil
	p.state = AVAILABLE
	return nil
}

func (p *ParkingSpot) isCompatible(vehicleType vehicle.VehicleType) bool {
	switch p.parkingSpotType {
	case COMPACT:
		return vehicleType == vehicle.CAR
	case LARGE:
		return vehicleType == vehicle.TRUCK || vehicleType == vehicle.CAR
	case TWOWHEELER:
		return vehicleType == vehicle.BIKE
	case HANDICAP:
		return vehicleType == vehicle.CAR
	default:
		return false
	}
}
