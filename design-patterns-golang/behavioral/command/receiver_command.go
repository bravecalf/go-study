package command

import "fmt"

// ReceiverCommand 具体command
type ReceiverCommand struct {
	receiver *Receiver
}

func (r *ReceiverCommand) execute() {
	if r.receiver == nil {
		fmt.Println("Please init receiver by setReceiver method.")
		return
	}
	r.receiver.action()
}

func (r *ReceiverCommand) setReceiver(receiver *Receiver) {
	r.receiver = receiver
}
