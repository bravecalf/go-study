package state

type Object struct {
	state State
	name  string
}

func (o *Object) setState(s State) {
	o.state = s
}

func (o *Object) execute() {
	o.state.doAction(o.name)
}
