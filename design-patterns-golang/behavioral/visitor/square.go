package visitor

type Square struct {
	side float64
}

func (s *Square) accept(v Visitor) {
	v.visitorForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}
