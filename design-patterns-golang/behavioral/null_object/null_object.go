package null_object

type NullObject struct {
}

func (n NullObject) getName() string {
	return "null object has no name"
}

func (n NullObject) isNil() bool {
	return true
}
