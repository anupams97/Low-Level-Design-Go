package vehicle

type VehicleType int

const (
	BIKE VehicleType = iota
	CAR
	TRUCK
)

type Vehicle interface {
	GetLicensePlate() string
	GetVehicleType() VehicleType
}
