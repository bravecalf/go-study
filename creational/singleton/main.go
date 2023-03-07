package singleton

import "sync"

func TestSingletonDemo() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				_ = getSingleInstance()
			}
		}()
	}
	wg.Wait()
}
