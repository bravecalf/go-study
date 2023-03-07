package builder

import (
	"errors"
	"fmt"
)

// WindowsBuilder windows生成器
type WindowsBuilder struct {
	Computer
}

func (w *WindowsBuilder) setCpu(cpu int) {
	w.Cpu = cpu
}

func (w *WindowsBuilder) setGpuType(gpuType string) {
	w.GpuType = gpuType
}

func (w *WindowsBuilder) setMemory(memory int) {
	w.Memory = memory
}

func (w *WindowsBuilder) getComputer() (*Computer, error) {
	if w.Cpu == 0 {
		fmt.Println("windows computer cpu is not set, do setCpu.")
		return nil, errors.New("computer cpu is not set")
	}

	if w.Memory == 0 {
		fmt.Println("windows computer memory is not set, do setMemory.")
		return nil, errors.New("computer memory is not set")
	}

	computer := &Computer{
		System:  "windows",
		Cpu:     w.Cpu,
		GpuType: w.GpuType,
		Memory:  w.Memory,
	}

	return computer, nil
}
