package command

import "fmt"

// Invoker 命令的调用者
type Invoker struct {
	command Command
}

func (i *Invoker) invoke() {
	if i.command == nil {
		fmt.Println("Please init command by setCommand method.")
		return
	}
	i.command.execute()
	return
}

func (i *Invoker) setCommand(cd Command) {
	i.command = cd
}

