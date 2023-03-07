package abstract_factory

import "fmt"

// ISportFactory 抽象工厂
type ISportFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportFactory, error) {
	if brand == "adidas" {
		return &AdidasFactory{}, nil
	}

	if brand == "nike" {
		return &NikeFactory{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}
