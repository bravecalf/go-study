package strategy

type OperationAdd struct {
}

func (o OperationAdd) doOperation(num1, num2 int) int {
	return num1 + num2
}
