package adapter

// Client 客户端
type Client struct {
}

// Animal 客户端接口
type Animal interface {
	walk()
}

// simulationAnimal  客户端调用，模拟动物行走
func (c Client) simulationAnimal(animal Animal) {
	animal.walk()
}
