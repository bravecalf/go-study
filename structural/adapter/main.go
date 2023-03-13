package adapter

func TestAdapterDemo() {
	client := Client{}
	dog := Dog{}
	client.simulationAnimal(dog)

	bird := Bird{}
	birdAdapter := BirdAdapter{bird: bird}
	client.simulationAnimal(birdAdapter)
}
