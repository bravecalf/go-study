package visitor

import (
	"fmt"
	"math"
)

type AreaVisitor struct {
	area float64
}

func (a *AreaVisitor) visitorForSquare(square *Square) {
	a.area = square.side * square.side
	fmt.Printf("calculate the area of square with side[%f] is [%f].\n", square.side, a.area)
}

func (a *AreaVisitor) visitorForCircle(circle *Circle) {
	a.area = math.Pi * circle.radius * circle.radius
	fmt.Printf("calculate the area of circle with redius[%f] is [%f].\n", circle.radius, a.area)
}
