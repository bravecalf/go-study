package mediator

import (
	"fmt"
	"time"
)

type Mediator struct {
	cache map[string]*User
}

func (m *Mediator) transferMessage(from, to, message string) {
	if user, ok := m.cache[to]; ok {
		user.receiverMessage = fmt.Sprintf("[%v][From %s] %s", time.Now(), from, message)
		fmt.Printf("send message from %s to %s successfully.\n", from, to)
		return
	} else {
		fmt.Println("not found user: ", to)
	}
	return
}
