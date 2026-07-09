package disk

import (
	"github.com/shirou/gopsutil/v3/disk"
)

type Disk struct {
	Total       uint64
	Used        uint64
	UsedPercent float64
}

func GetDisk(path string) (Disk, error) {

	u, err := disk.Usage(path)
	if err != nil {
		return Disk{}, err
	}

	return Disk{
		Total:       u.Total,
		Used:        u.Used,
		UsedPercent: u.UsedPercent,
	}, nil
}