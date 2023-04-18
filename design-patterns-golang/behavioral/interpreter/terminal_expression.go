package interpreter

import "strings"

type TerminalExpression struct {
	data string
}

func (t TerminalExpression) interpret(context string) bool {
	return strings.Contains(context, t.data)
}
