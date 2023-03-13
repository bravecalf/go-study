package adapter

import "fmt"

type Bird struct {
}

func (b Bird) fly() {
	fmt.Println("bird is flying ...")
}
