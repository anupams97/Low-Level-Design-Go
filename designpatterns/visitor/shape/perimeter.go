package shape

type Perimeter struct {
	Perimeter float64
}

func (p *Perimeter) visitCircle(circle *Circle) {
	p.Perimeter = 2.0 * 3.14 * circle.radius
}

func (p *Perimeter) visitSquare(square *Square) {
	p.Perimeter = 4 * square.side
}
