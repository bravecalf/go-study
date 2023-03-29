package flyweight

import "math/rand"

func TestFlyweightDemo() {
	colors := []string{"Red", "Green", "Black", "Yellow", "White"}
	getRandomColor := func() string {
		return colors[rand.Intn(len(colors))]
	}
	for i := 0; i < 10; i++ {
		circle := getCircle(getRandomColor())
		circle.setX(rand.Intn(1000))
		circle.setY(rand.Intn(1000))
		circle.setRadius(100)
		circle.draw()
	}
}
