package bridge

import "fmt"

type Canon struct {
}

func (c Canon) printFile() {
	fmt.Println("Canon printer is working ...")
}
