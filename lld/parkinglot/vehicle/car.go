package vehicle

type Car struct {
	licensePlate string
}

func (c *Car) GetLicensePlate() string {
	return c.licensePlate
}

func (c *Car) GetVehicleType() VehicleType {
	return CAR
}

func NewCar(lp string) Vehicle {
	return &Car{licensePlate: lp}
}
