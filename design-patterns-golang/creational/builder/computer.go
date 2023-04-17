package builder

import "encoding/json"

// Computer 具体产品
type Computer struct {
	System  string `json:"system"`
	Cpu     int    `json:"cpu"`
	GpuType string `json:"gpuType,omitempty"`
	Memory  int    `json:"memory"`
}

func (c *Computer) printParameters() string {
	computer, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(computer)
}
