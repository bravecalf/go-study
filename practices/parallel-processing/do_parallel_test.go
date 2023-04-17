package parallel_processing

import (
	"fmt"
	"sync"
	"testing"
)

func TestDoParallelProcessing(t *testing.T) {
	result := make([]int, 0)
	mutex := sync.Mutex{}
	f := func(a interface{}) error {
		x := a.(int)
		mutex.Lock()
		result = append(result, x)
		mutex.Unlock()
		return nil
	}
	productions := make([]int, 100)
	if err := DoParallelProcessing(20, f, productions); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(result)
}
