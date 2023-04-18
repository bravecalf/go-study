package command

import "fmt"

// Receiver 命令的接收者
type Receiver struct {
}

func (r *Receiver) action() {
	fmt.Println("receiver is processing...")
}
