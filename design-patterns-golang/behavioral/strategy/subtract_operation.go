package strategy

type OperationSubtract struct {
}

func (o OperationSubtract) doOperation(num1, num2 int) int {
	return num1 - num2
}
