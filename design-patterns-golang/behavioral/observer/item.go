package observer

import "fmt"

type Item struct {
	observerMap map[string]Observer
	name        string
}

func (i *Item) register(observer Observer) {
	if _, ok := i.observerMap[observer.getID()]; ok {
		return
	}
	i.observerMap[observer.getID()] = observer
	return
}

func (i *Item) deregister(observer Observer) {
	fmt.Printf("deregister observer[%s] with item[%s].\n", observer.getID(), i.name)
	delete(i.observerMap, observer.getID())
}

func (i *Item) notifyAll() {
	for _, observer := range i.observerMap {
		observer.update(i.name)
	}
}
