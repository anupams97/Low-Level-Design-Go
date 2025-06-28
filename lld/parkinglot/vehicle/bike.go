package vehicle

type Bike struct {
	licensePlate string
}

func (b *Bike) GetLicensePlate() string {
	return b.licensePlate
}

func (b *Bike) GetVehicleType() VehicleType {
	return BIKE
}

func NewBike(lp string) Vehicle {
	return &Bike{licensePlate: lp}
}
