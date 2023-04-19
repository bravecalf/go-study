package state

func TestStateDemo() {
	// 创建state
	onState := OnState{}
	offState := OffState{}

	// 创建对象
	tv := &Object{
		name:     "Tv",
		onState:  onState,
		offState: offState,
	}

	// 变更对象状态
	tv.setState(onState)
	tv.execute()
	tv.execute()
	tv.execute()
}
