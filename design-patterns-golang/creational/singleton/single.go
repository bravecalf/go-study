package singleton

import (
	"fmt"
	"sync"
)

type single struct {
}

var lock = sync.Mutex{}

var singleInstance *single

func getSingleInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("create single instance now.")
			singleInstance = new(single)
		} else {
			fmt.Println("single instance is already created-1.")
		}
	} else {
		fmt.Println("single instance is already created-2.")
	}
	return singleInstance
}
