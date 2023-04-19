package state

import "fmt"

// OffState 具体状态
type OffState struct {
}

func (o OffState) doAction(object *Object) {
	fmt.Printf("turn %s off\n", object.name)
	object.state = object.onState
}
