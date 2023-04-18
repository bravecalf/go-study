package command

func TestCommandDemo() {
	// 创建接收者
	receiver := new(Receiver)
	// 创建command
	command := new(ReceiverCommand)
	// 设置command的接收者
	command.setReceiver(receiver)
	// 创建调用者
	invoker := new(Invoker)
	// 设置调用者的command
	invoker.setCommand(command)
	// 调用者调用命令
	invoker.invoke()
}
