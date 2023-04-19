package memento

import (
	"errors"
)

// CareTaker 管理者
type CareTaker struct {
	mementoList []*Memento
}

func (c *CareTaker) addMemento(m *Memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *CareTaker) getMemento(index int) (*Memento, error) {
	if index > len(c.mementoList) {
		return nil, errors.New("index exceeds array bounds")
	}
	return c.mementoList[index], nil
}
