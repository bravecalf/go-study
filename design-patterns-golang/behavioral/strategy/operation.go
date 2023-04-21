package strategy

type Operation interface {
	doOperation(int, int) int
}
