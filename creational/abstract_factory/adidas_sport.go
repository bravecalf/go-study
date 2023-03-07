package abstract_factory

import "fmt"

// AdidasShoe 具体产品
type AdidasShoe struct {
}

func (n AdidasShoe) getShoeDetail() string {
	return fmt.Sprintf("adidas shoe.")
}

// AdidasShirt 具体产品
type AdidasShirt struct {
}

func (n AdidasShirt) getShirtDetail() string {
	return fmt.Sprintf("adidas shirt.")
}
