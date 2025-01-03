package lot

import (
	"github.com/anupams97/Low-Level-Design-Go/parkinglot/vehicle"
)

type ParkingStrategy interface {
	AllocateSpot([]*ParkingLevel, vehicle.VehicleType) *ParkingSpot
}
