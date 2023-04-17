package flyweight

var CircleColorMap = make(map[string]*Circle)

func getCircle(color string) *Circle {
	if circle, ok := CircleColorMap[color]; ok {
		return circle
	}
	circle := &Circle{color: color}
	CircleColorMap[color] = circle
	return circle
}
