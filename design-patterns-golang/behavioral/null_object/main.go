package null_object

import "fmt"

func TestNullObjectDemo() {
	// 创建对象工厂
	objectFactory := ObjectFactory{
		nameList: []string{"A", "B", "D"},
	}

	// 获取具体对象
	o1 := objectFactory.getCustomObject("A")
	o2 := objectFactory.getCustomObject("B")
	o3 := objectFactory.getCustomObject("C")
	o4 := objectFactory.getCustomObject("D")

	fmt.Printf("o1 is exist? [%v], name is [%s]\n", !o1.isNil(), o1.getName())
	fmt.Printf("o2 is exist? [%v], name is [%s]\n", !o2.isNil(), o2.getName())
	fmt.Printf("o3 is exist? [%v], name is [%s]\n", !o3.isNil(), o3.getName())
	fmt.Printf("o4 is exist? [%v], name is [%s]\n", !o4.isNil(), o4.getName())
}
