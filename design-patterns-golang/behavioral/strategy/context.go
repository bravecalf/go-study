package strategy

type MyContext struct {
	operation Operation
}

func (m *MyContext) setOperation(operation Operation) {
	m.operation = operation
}

func (m *MyContext) executeOperation(num1, num2 int) int {
	return m.operation.doOperation(num1, num2)
}
