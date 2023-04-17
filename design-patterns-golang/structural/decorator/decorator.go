package decorator

// Decorator 抽象装饰
type Decorator interface {
	Component
	setComponent(component Component)
	addedFunction()
}
