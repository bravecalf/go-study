package builder

import "fmt"

func TestBuilderDemo() {
	// 构建生成器
	winBuilder, err := getBuilder("windows")
	if err != nil {
		panic(err)
	}
	// 构建主管
	wDirector := newDirector(winBuilder)
	// 参数配置
	wDirector.builder.setCpu(4)
	wDirector.builder.setMemory(16)

	// 生成产品
	winComputer := wDirector.buildComputer()
	fmt.Println(winComputer.printParameters())

	linuxBuilder, err := getBuilder("linux")
	if err != nil {
		panic(err)
	}
	lDirector := newDirector(linuxBuilder)
	lDirector.builder.setCpu(8)
	lDirector.builder.setMemory(32)
	lDirector.builder.setGpuType("A100")

	linuxComputer := lDirector.buildComputer()
	fmt.Println(linuxComputer.printParameters())

}
