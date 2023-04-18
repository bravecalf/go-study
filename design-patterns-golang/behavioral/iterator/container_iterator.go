package iterator

type ContainerIterator struct {
	index      int
	length     int
	containers []Container
}

func (c *ContainerIterator) hasNext() bool {
	return c.index < c.length
}

func (c *ContainerIterator) getNext() interface{} {
	if c.hasNext() {
		container := c.containers[c.index]
		c.index++
		return container
	}
	return nil
}
