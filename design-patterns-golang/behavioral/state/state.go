package state

// State 状态接口
type State interface {
	doAction(*Object)
}
