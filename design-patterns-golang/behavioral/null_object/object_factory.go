package null_object

type ObjectFactory struct {
	nameList []string
}

func (o ObjectFactory) getCustomObject(name string) CustomObject {
	for _, itemName := range o.nameList {
		if itemName == name {
			return RealObject{name: name}
		}
	}
	return NullObject{}
}
