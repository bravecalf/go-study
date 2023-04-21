package visitor

type Circle struct {
	radius float64
}

func (c *Circle) accept(v Visitor) {
	v.visitorForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}
