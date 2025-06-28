package elevatorsystem

import (
	"time"
)

type Direction int

func (d Direction) String() string {
	return [...]string{"UP", "DOWN"}[d]
}

const (
	UP Direction = iota
	DOWN
)

type ElevatorCar struct {
	id           int
	capacity     int
	request      chan int
	stop         chan struct{}
	currentFloor int
	direction    Direction
}

func NewElevatorCar(id, capacity, cf int, direction Direction) *ElevatorCar {
	return &ElevatorCar{
		id:           id,
		capacity:     capacity,
		request:      make(chan int),
		stop:         make(chan struct{}),
		currentFloor: cf,
		direction:    direction,
	}
}

func (e *ElevatorCar) AddRequests(req int) {
	e.request <- req
}

func (e *ElevatorCar) Run() {
	go func() {
		for {
			select {
			case request := <-e.request:
				e.processRequest(request)
			case <-e.stop:
				return

			}
		}
	}()
}

func (e *ElevatorCar) setElevatorFloor(floor int) {
	e.currentFloor = floor
}
func (e *ElevatorCar) getCurrentFloor() int {
	return e.currentFloor
}

func (e *ElevatorCar) processRequest(floor int) {
	if e.getCurrentFloor() > floor {
		e.direction = DOWN
	} else {
		e.direction = UP
	}
	time.Sleep(1 * time.Second)
	e.setElevatorFloor(floor)
}

func (e *ElevatorCar) Stop() {
	close(e.stop)
}
