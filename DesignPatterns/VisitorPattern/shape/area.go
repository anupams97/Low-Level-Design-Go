package shape

import "math"

type Area struct {
	Area float64
}

func (a *Area) visitCircle(shape *Circle) {
	a.Area = math.Pow(shape.radius, 2) * 3.14
}
func (a *Area) visitSquare(shape *Square) {
	a.Area = math.Pow(shape.side, 2)
}
