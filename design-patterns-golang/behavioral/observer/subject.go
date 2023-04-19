package observer

// Subject 实体接口
type Subject interface {
	register(Observer)
	deregister(Observer)
	notifyAll()
}
