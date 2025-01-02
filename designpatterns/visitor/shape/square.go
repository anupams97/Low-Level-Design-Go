package shape

type Square struct {
	side float64
}

func (s *Square) SetSide(side float64) {
	s.side = side
}

func (s *Square) Accept(v Visitor) {
	v.visitSquare(s)
}
