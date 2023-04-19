package state

import "fmt"

// OffState 具体状态
type OffState struct {
}

func (o OffState) doAction(object string) {
	fmt.Printf("turn %s off\n", object)
}
