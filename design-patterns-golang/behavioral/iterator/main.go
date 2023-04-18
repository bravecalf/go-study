package iterator

import "fmt"

func TestIteratorDemo() {
	container1 := Container{id: 2021}
	container2 := Container{id: 2024}
	containers := []Container{container1, container2}

	containerIterator := &ContainerIterator{
		length:     len(containers),
		containers: containers,
	}

	for containerIterator.hasNext() {
		item := containerIterator.getNext().(Container)
		fmt.Printf("container: %#v\n", item)
	}
}
