package bridge

import "fmt"

type HP struct {
}

func (h HP) printFile() {
	fmt.Println("HP printer is working ...")
}
