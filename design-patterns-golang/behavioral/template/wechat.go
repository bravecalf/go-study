package template

import "fmt"

type Wechat struct {
}

func (w Wechat) login(id string) {
	fmt.Printf("Welcome to wechat %s, have fun!\n", id)
}

func (w Wechat) play() {
	fmt.Println("chatting with friends ...")
}

func (w Wechat) exit(id string) {
	fmt.Printf("GoodBye %s, Welcome back again.\n", id)
}
