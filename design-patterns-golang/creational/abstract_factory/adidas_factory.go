package abstract_factory

// AdidasFactory 具体工厂
type AdidasFactory struct {
}

func (a *AdidasFactory) makeShoe() IShoe {
	return AdidasShoe{}
}

func (a *AdidasFactory) makeShirt() IShirt {
	return AdidasShirt{}
}
