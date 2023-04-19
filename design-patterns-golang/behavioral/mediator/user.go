package mediator

type User struct {
	name            string
	receiverMessage string
	chatRoom        *Mediator
}

func (u *User) sendMessage(name, message string) {
	u.chatRoom.transferMessage(u.name, name, message)
	return
}

func (u *User) addChatRoom(mediator *Mediator) {
	mediator.cache[u.name] = u
	u.chatRoom = mediator
}
