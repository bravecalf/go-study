package abstract_factory

import "fmt"

// NikeShoe 具体产品
type NikeShoe struct {
}

func (n NikeShoe) getShoeDetail() string {
	return fmt.Sprintf("nike shoe.")
}

// NikeShirt 具体产品
type NikeShirt struct {
}

func (n NikeShirt) getShirtDetail() string {
	return fmt.Sprintf("nike shirt.")
}
