package cpu

import (
	"time"
	"github.com/shirou/gopsutil/v3/cpu"
)

type CPU struct {
	Usage float64
	Cores int
	Model string
}

func GetCPU() (CPU, error) {

	// Get usage percentage
	usage, err := cpu.Percent(time.Second, false)

	if err != nil {
		return CPU{}, err
	}

	// Get number of cores
	cores, err := cpu.Counts(true)

	if err != nil {
		return CPU{}, err
	}

	// Get model information
	info, err := cpu.Info()

	if err != nil {
		return CPU{}, err
	}

	model := ""

	if len(info) > 0 {
		model = info[0].ModelName
	}

	return CPU{
		Usage: usage[0],
		Cores: cores,
		Model: model,
	}, nil
}