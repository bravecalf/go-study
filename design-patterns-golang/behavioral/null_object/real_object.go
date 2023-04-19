package null_object

type RealObject struct {
	name string
}

func (r RealObject) getName() string {
	return r.name
}

func (r RealObject) isNil() bool {
	return false
}
