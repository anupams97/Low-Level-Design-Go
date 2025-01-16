package app

import "time"

type Movie struct {
	ID   int
	Name string
}

type Theater struct {
}

type Location struct {
}

type User struct {
}

type Show struct {
	ID        int
	StartTime time.Time
	EndTime   time.Time
	Screen    Screen
}
type Screen struct {
	Seats []Seat
}

type Seat struct {
	ID   int
	Type SeatType
	Cost int
}
