package lot

type ParkingLevel struct {
	level        int
	parkingSpots []*ParkingSpot
}

func NewParkingLevel(level int, levelConfig map[ParkingSpotType]int) *ParkingLevel {
	l := &ParkingLevel{level: level}
	spotID := 1
	for spotType, count := range levelConfig {
		for i := 0; i < count; i++ {
			l.parkingSpots = append(l.parkingSpots, NewParkingSpot(spotID, level, spotType))
			spotID += 1
		}
	}
	return l
}
