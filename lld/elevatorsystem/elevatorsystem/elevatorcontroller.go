package elevatorsystem

import "sync"

var once sync.Once
var instance *ElevatorController

type ElevatorController struct {
	optimalElevatorStrategy Strategy
	elevators               []*ElevatorCar
	requests                chan *Request
	stop                    chan struct{}
}

func NewElevatorController(optimalElevatorStrategy Strategy, noElevator int) *ElevatorController {
	once.Do(func() {
		var elevatorCars []*ElevatorCar
		for i := 0; i < noElevator; i++ {
			ec := NewElevatorCar(i, 10, 0, UP)
			elevatorCars = append(elevatorCars, ec)
			ec.Run()
		}
		instance = &ElevatorController{optimalElevatorStrategy: optimalElevatorStrategy, elevators: elevatorCars}
		instance.run()
	})
	return instance
}

func (ec *ElevatorController) AddRequests(req *Request) {
	ec.requests <- req
}

func (ec *ElevatorController) run() {
	go func() {
		for {
			select {
			case req := <-ec.requests:
				ec.processRequest(req)
			case <-ec.stop:
				ec.Stop()
				return
			}
		}
	}()

}

func (ec *ElevatorController) processRequest(req *Request) {

}

type Strategy interface {
	getOptimalElevator() *ElevatorCar
}

func (ec *ElevatorController) Stop() {
	//stopping all elevators
	for _, c := range ec.elevators {
		c.Stop()
	}
	//stopping the elevator controller
	close(ec.stop)
}
