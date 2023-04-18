package interpreter

type OrExpression struct {
	exp1 Expression
	exp2 Expression
}

func (o OrExpression) interpret(context string) bool {
	return o.exp1.interpret(context) || o.exp2.interpret(context)
}
