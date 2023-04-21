package strategy

import "fmt"

func TestStrategyDemo() {
	// 创建上下文
	ct := &MyContext{}

	// 创建操作策略
	addOp := OperationAdd{}
	subOp := OperationSubtract{}

	// 修改上下文策略对象,执行上下文策略
	ct.setOperation(addOp)
	fmt.Println("1 + 2 =", ct.executeOperation(1, 2))

	ct.setOperation(subOp)
	fmt.Println("2 - 1 =", ct.executeOperation(2, 1))

}
