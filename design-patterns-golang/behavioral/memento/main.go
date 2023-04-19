package memento

import "fmt"

func TestMementoDemo() {
	//创建管理者
	careTaker := &CareTaker{
		mementoList: make([]*Memento, 0),
	}

	// 创建原发器
	originator := &Originator{state: "A"}
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	careTaker.addMemento(originator.createMemento())

	originator.setState("B")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	careTaker.addMemento(originator.createMemento())

	originator.setState("C")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	careTaker.addMemento(originator.createMemento())

	// 回滚状态到索引0
	mementoIndex0, _ := careTaker.getMemento(0)
	originator.restoreMemento(mementoIndex0)
	fmt.Printf("Originator Current State: %s\n", originator.getState())

	// 回滚状态到索引1
	mementoIndex1, _ := careTaker.getMemento(1)
	originator.restoreMemento(mementoIndex1)
	fmt.Printf("Originator Current State: %s\n", originator.getState())
}
