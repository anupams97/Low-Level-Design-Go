package car

// director
type Director struct {
	CarBuilder
}

func NewDirector(cb CarBuilder) *Director {
	return &Director{cb}
}

func (d *Director) CreateSportsCar(name string, hp float64) Car {
	return d.CarBuilder.setName(name).setHP(hp).build()
}
func (d *Director) CreateOffRoadCar(name string, hp float64) Car {
	return d.CarBuilder.setName(name).setHP(hp).build()
}

// buider interface
type CarBuilder interface {
	setName(name string) CarBuilder
	setHP(hp float64) CarBuilder
	build() Car
}

// sports  car  builder that implents builder interface
type sportsCarBuilder struct {
	Car
}

func NewSportsCarBuider() *sportsCarBuilder {
	return &sportsCarBuilder{}
}
func (s *sportsCarBuilder) setName(name string) CarBuilder {
	s.Car.Name = name
	return s
}

func (s *sportsCarBuilder) setHP(hp float64) CarBuilder {
	s.Car.HP = hp
	return s
}

func (o *sportsCarBuilder) build() Car {
	return o.Car
}

// offroad car builder that implements builder interface
type offRoadCarBuilder struct {
	Car
}

func NewOffroadCarBuider() *offRoadCarBuilder {
	return &offRoadCarBuilder{}
}

func (o *offRoadCarBuilder) setName(name string) CarBuilder {
	o.Car.Name = name
	return o
}

func (o *offRoadCarBuilder) setHP(hp float64) CarBuilder {
	o.Car.HP = hp
	return o
}

func (o *offRoadCarBuilder) build() Car {
	return o.Car
}

// Car  struct
type Car struct {
	Name string
	HP   float64
}
