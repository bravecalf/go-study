package interpreter

type Expression interface {
	interpret(string) bool
}
