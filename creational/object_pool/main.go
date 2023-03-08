package object_pool

import (
	"fmt"
	"strconv"
)

func TestObjectPoolDemo() {
	// 创建具体连接
	connections := make([]IPoolObject, 0)
	for i := 0; i < 5; i++ {
		cn := &Connection{ID: strconv.Itoa(i)}
		connections = append(connections, cn)
	}

	// 初始化连接池
	pool, err := initPool(connections)
	if err != nil {
		panic(err)
	}

	// 调用连接
	con1, err := pool.loan()
	if err == nil {
		fmt.Printf("==== use connection [%s] from pool.\n", con1.getID())
	}

	con2, err := pool.loan()
	if err == nil {
		fmt.Printf("==== use connection [%s] from pool.\n", con2.getID())
	}

	fmt.Printf("The capacity of pool: %d\n", pool.GetCapacity())
	fmt.Printf("The len of pool: %d\n", pool.Len())

	// 释放连接
	err = pool.receive(con1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The capacity of pool: %d\n", pool.GetCapacity())
	fmt.Printf("The len of pool: %d\n", pool.Len())
}
