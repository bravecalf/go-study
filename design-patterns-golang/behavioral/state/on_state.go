package state

import "fmt"

// OnState 具体状态
type OnState struct {
}

func (o OnState) doAction(object *Object) {
	fmt.Printf("turn %s on\n", object.name)
	object.state = object.offState
}
