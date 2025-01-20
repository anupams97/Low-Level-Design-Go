package app

import "time"

type BookingSystem struct {
	Users      []User
	Controller []TheaterController
}

type TheaterController struct {
}

type Movie struct {
	ID   int
	Name string
}

type Theater struct {
	ID      string
	Shows   []Show
	Screens []Screen
}

type Location struct {
	ID   string
	City string
}

type User struct {
	Name     string
	Location Location
}

type Show struct {
	ID        int
	StartTime time.Time
	EndTime   time.Time
	Screen    Screen
	Movie     Movie
}
type Screen struct {
	Seats []Seat
}

type Seat struct {
	ID   int
	Type SeatType
	Cost int
}

/*
Movie booking system
1. Search for a movie
2. Book that movie on a particular show in a particular theater
3. make sure concurrent seat booking works well


*/
