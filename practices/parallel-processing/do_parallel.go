package parallel_processing

import (
	"reflect"
	"sync"
)

func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}

// DoParallelProcessing 并发处理对象列表
func DoParallelProcessing(workersNum int, doPart func(interface{}) error, args interface{}) error {
	workChan := make(chan struct{}, workersNum)
	wg := sync.WaitGroup{}
	var err error
	productions := ToSlice(args)
	for _, production := range productions {
		workChan <- struct{}{}
		wg.Add(1)
		go func(item interface{}) {
			defer func() {
				<-workChan
				wg.Done()
			}()
			if partErr := doPart(item); partErr != nil {
				err = partErr
			}
		}(production)

		if err != nil {
			break
		}
	}

	wg.Wait()
	return err
}
