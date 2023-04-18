package interpreter

type AndExpression struct {
	exp1 Expression
	exp2 Expression
}

func (a AndExpression) interpret(context string) bool {
	return a.exp1.interpret(context) && a.exp2.interpret(context)
}
