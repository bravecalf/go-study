package builder

import (
	"errors"
	"fmt"
)

// LinuxBuilder Linux生成器
type LinuxBuilder struct {
	Computer
}

func (l *LinuxBuilder) setCpu(cpu int) {
	l.Cpu = cpu
}

func (l *LinuxBuilder) setGpuType(gpuType string) {
	l.GpuType = gpuType
}

func (l *LinuxBuilder) setMemory(memory int) {
	l.Memory = memory
}

func (l *LinuxBuilder) getComputer() (*Computer, error) {
	if l.Cpu == 0 {
		fmt.Println("linux computer cpu is not set, do setCpu.")
		return nil, errors.New("computer cpu is not set")
	}

	if l.Memory == 0 {
		fmt.Println("linux computer memory is not set, do setMemory.")
		return nil, errors.New("computer memory is not set")
	}

	computer := &Computer{
		System:  "Linux",
		Cpu:     l.Cpu,
		GpuType: l.GpuType,
		Memory:  l.Memory,
	}

	return computer, nil
}
