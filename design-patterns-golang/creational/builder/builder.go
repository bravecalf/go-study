package builder

import "errors"

// IBuilder 生成器接口
type IBuilder interface {
	setCpu(cpu int)
	setGpuType(gpuType string)
	setMemory(memory int)
	getComputer() (*Computer, error)
}

func getBuilder(system string) (IBuilder, error) {
	switch system {
	case "windows":
		return &WindowsBuilder{}, nil
	case "linux":
		return &LinuxBuilder{}, nil
	default:
		return nil, errors.New("wrong sys type passed")
	}
}
