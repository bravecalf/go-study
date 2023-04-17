package adapter

import "fmt"

type Dog struct {
}

func (d Dog) walk() {
	fmt.Println("dog is walking ...")
}
