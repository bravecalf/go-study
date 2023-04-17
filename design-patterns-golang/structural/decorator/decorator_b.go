package decorator

import "fmt"

// DecoratorB 具体装饰
type DecoratorB struct {
	component Component
}

func (d *DecoratorB) operation() {
	d.component.operation()
	d.addedFunction()
}

func (d *DecoratorB) setComponent(component Component) {
	d.component = component
}

func (d *DecoratorB) addedFunction() {
	fmt.Println("add function from DecoratorB")
}
