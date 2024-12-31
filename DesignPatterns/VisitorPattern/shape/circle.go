package shape

type Circle struct {
	radius float64
}

func (s *Circle) SetRadius(radius float64) {
	s.radius = radius
}

func (c *Circle) Accept(v Visitor) {
	v.visitCircle(c)
}
