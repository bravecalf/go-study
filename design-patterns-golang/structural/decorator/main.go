package decorator

func TestDecoratorDemo() {
	c := &ActualComponent{}
	c.operation()

	d1 := &DecoratorA{c}
	d1.operation()

	d2 := &DecoratorB{d1}
	d2.operation()
}
