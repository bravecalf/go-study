package decorator

import "fmt"

// Component 抽象构件
type Component interface {
	operation()
}

// ActualComponent 具体构件
type ActualComponent struct {
}

func (a *ActualComponent) operation() {
	fmt.Println("actual component operation.")
}
