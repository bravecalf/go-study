package decorator

import "fmt"

// DecoratorA 具体装饰
type DecoratorA struct {
	component Component
}

func (d *DecoratorA) operation() {
	d.component.operation()
	d.addedFunction()
}

func (d *DecoratorA) setComponent(component Component) {
	d.component = component
}

func (d *DecoratorA) addedFunction() {
	fmt.Println("add function from decoratorA")
}
