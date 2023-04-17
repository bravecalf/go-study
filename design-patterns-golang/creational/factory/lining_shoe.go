package factory

import "errors"

const LiNingLogo = "lining"

// LiNingShoe lining产品
type LiNingShoe struct {
	color string
	size  int
}

func (n *LiNingShoe) getLogo() string {
	return LiNingLogo
}

func (n *LiNingShoe) setColor(color string) {
	n.color = color
}

func (n *LiNingShoe) getColor() (string, error) {
	if n.color == "" {
		return "", errors.New("color is empty")
	}
	return n.color, nil
}

func (n *LiNingShoe) setSize(size int) {
	n.size = size
}

func (n *LiNingShoe) getSize() (int, error) {
	if n.size == 0 {
		return 0, errors.New("size is zero")
	}
	return n.size, nil
}

func newLiNingShoe(color string, size int) *LiNingShoe {
	return &LiNingShoe{
		color: color,
		size:  size,
	}
}
