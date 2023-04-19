package mediator

import "fmt"

func TestMediatorDemo() {
	// 初始化mediator
	chatRoom := &Mediator{
		cache: make(map[string]*User),
	}

	// 创建user对象
	li := &User{name: "Li"}
	zhang := &User{name: "Zhang"}
	li.addChatRoom(chatRoom)
	zhang.addChatRoom(chatRoom)

	// li 发送信息给 zhang
	li.sendMessage("Zhang", "Hello, zhang.")

	fmt.Println(zhang.receiverMessage)
}
