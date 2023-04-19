package observer

func TestObserverDemo() {
	// 创建具体商品
	production := &Item{
		observerMap: make(map[string]Observer),
		name:        "iphone 15",
	}

	// 创建多个observer用户
	customer1 := &Customer{id: "xy@baidu.com"}
	customer2 := &Customer{id: "wx@gamil.com"}

	// 将用户与产品进行关联
	production.register(customer1)
	production.register(customer2)

	production.notifyAll()

	// 将用户2与产品进行解绑
	production.deregister(customer2)
	production.notifyAll()

}
