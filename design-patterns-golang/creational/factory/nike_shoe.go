package factory

import "errors"

const NikeLogo = "nike"

// NikeShoe nike产品
type NikeShoe struct {
	color string
	size  int
}

func (n *NikeShoe) getLogo() string {
	return NikeLogo
}

func (n *NikeShoe) setColor(color string) {
	n.color = color
}

func (n *NikeShoe) getColor() (string, error) {
	if n.color == "" {
		return "", errors.New("color is empty")
	}
	return n.color, nil
}

func (n *NikeShoe) setSize(size int) {
	n.size = size
}

func (n *NikeShoe) getSize() (int, error) {
	if n.size == 0 {
		return 0, errors.New("size is zero")
	}
	return n.size, nil
}

func newNikeShoe(color string, size int) *NikeShoe {
	return &NikeShoe{
		color: color,
		size:  size,
	}
}
