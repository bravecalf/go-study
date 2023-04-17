package adapter

import "fmt"

type BirdAdapter struct {
	bird Bird
}

func (b BirdAdapter) walk() {
	fmt.Println("Adapter convert walk to fly.")
	b.bird.fly()
}
