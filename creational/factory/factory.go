package factory

import "errors"

// ShoeFactory 产品工厂
type ShoeFactory struct {
	Color string
	Size  int
}

func (s *ShoeFactory) GetShoe(shoeType string) (IShoe, error) {
	switch shoeType {
	case NikeLogo:
		return newNikeShoe(s.Color, s.Size), nil
	case LiNingLogo:
		return newLiNingShoe(s.Color, s.Size), nil
	default:
		return nil, errors.New("unknown shoe type")
	}
}

func NewShoeFactory(color string, size int) *ShoeFactory {
	return &ShoeFactory{
		Color: color,
		Size:  size,
	}
}
