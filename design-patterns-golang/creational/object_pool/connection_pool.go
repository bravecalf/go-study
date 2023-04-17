package object_pool

import (
	"fmt"
	"sync"
)

type Pool struct {
	idle     []IPoolObject
	active   []IPoolObject
	capacity int
	mulock   *sync.Mutex
}

// InitPool Initialize the pool
func initPool(poolObjects []IPoolObject) (*Pool, error) {
	if len(poolObjects) == 0 {
		return nil, fmt.Errorf("cannot craete a pool of 0 length")
	}
	pool := &Pool{
		idle:     poolObjects,
		active:   make([]IPoolObject, 0),
		capacity: len(poolObjects),
		mulock:   new(sync.Mutex),
	}

	return pool, nil
}

// Loan try to get connection from pool
func (p *Pool) loan() (IPoolObject, error) {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	if len(p.idle) == 0 {
		return nil, fmt.Errorf("no pool object free. please request after sometime")
	}
	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	fmt.Printf("Loan Pool Object with ID: %s\n", obj.getID())
	return obj, nil
}

// Receive move targetObject from active to idle in pool
func (p *Pool) receive(targetObject IPoolObject) error {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	err := p.remove(targetObject)
	if err != nil {
		return err
	}
	p.idle = append(p.idle, targetObject)
	fmt.Printf("receive Pool Object with ID: %s\n", targetObject.getID())
	return nil

}

// Remove release object connection from active of pool
func (p *Pool) remove(targetObject IPoolObject) error {
	curActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.getID() == targetObject.getID() {
			p.active[curActiveLength-1], p.active[i] = p.active[i], p.active[curActiveLength-1]
			p.active = p.active[:curActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("targe pool object doesn't belong to the pool")
}

// GetCapacity get the capacity of pool
func (p *Pool) GetCapacity() int {
	return p.capacity
}

// Len get current the number of idle connections
func (p *Pool) Len() int {
	return len(p.idle)
}
