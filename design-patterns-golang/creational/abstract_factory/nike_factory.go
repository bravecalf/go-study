package abstract_factory

// NikeFactory 具体工厂
type NikeFactory struct {
}

func (n *NikeFactory) makeShoe() IShoe {
	return NikeShoe{}
}

func (n *NikeFactory) makeShirt() IShirt {
	return NikeShirt{}
}
