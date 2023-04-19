package state

func TestStateDemo() {
	// 创建对象
	tv := &Object{name: "Tv"}

	// 创建state
	onState := OnState{}
	offState := OffState{}

	// 变更对象状态
	tv.setState(onState)
	tv.execute()

	tv.setState(offState)
	tv.execute()
}