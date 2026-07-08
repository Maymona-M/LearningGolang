package memory

import (
	"github.com/shirou/gopsutil/v3/mem"
)

type Memory struct {
	Total       uint64
	Used        uint64
	UsedPercent float64
}

func GetMemory() (Memory, error) {

	v, err := mem.VirtualMemory()
	if err != nil {
		return Memory{}, err
	}

	return Memory{
		Total:       v.Total,
		Used:        v.Used,
		UsedPercent: v.UsedPercent,
	}, nil
}