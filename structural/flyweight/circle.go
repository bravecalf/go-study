package flyweight

import "fmt"

type Circle struct {
	color        string
	x, y, radius int
}

func (c *Circle) setX(x int) {
	c.x = x
}

func (c *Circle) setY(y int) {
	c.y = y
}

func (c *Circle) setRadius(radius int) {
	c.radius = radius
}

func (c *Circle) draw() {
	fmt.Printf("circle draw [%#v] \n", c)
}
