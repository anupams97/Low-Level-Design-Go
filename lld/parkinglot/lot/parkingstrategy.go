package lot

import (
	"github.com/anupams97/Low-Level-Design-Go/lld/parkinglot/vehicle"
)

type ParkingStrategy interface {
	AllocateSpot([]*ParkingLevel, vehicle.VehicleType) *ParkingSpot
}
