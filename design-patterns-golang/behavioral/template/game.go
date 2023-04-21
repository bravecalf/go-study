package template

import "fmt"

type Game struct {
}

func (g Game) login(id string) {
	fmt.Printf("Welcome to Game %s, have fun!\n", id)
}

func (g Game) play() {
	fmt.Println("playing game ...")
}

func (g Game) exit(id string) {
	fmt.Printf("GoodBye %s, Welcome back again.\n", id)
}
