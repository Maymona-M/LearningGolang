package main

import (
	"fmt"
	"system-monitor/cpu"
	"system-monitor/disk"
	"system-monitor/memory"
)

func main() {

	cpuInfo, err := cpu.GetCPU()
	if err != nil {
		fmt.Println("Error retrieving CPU info:", err)
		return
	}

	memInfo, err := memory.GetMemory()
	if err != nil {
		fmt.Println("Error retrieving memory info:", err)
		return
	}

	diskInfo, err := disk.GetDisk("/")
	if err != nil {
		fmt.Println("Error retrieving disk info:", err)
		return
	}

	fmt.Println("===== System Information =====")

	const GB = 1024 * 1024 * 1024

	fmt.Println("CPU:")
	fmt.Printf("   Usage: %.1f%%\n", cpuInfo.Usage)
	fmt.Printf("   Cores: %d\n", cpuInfo.Cores)
	fmt.Printf("   Model: %s\n", cpuInfo.Model)

	fmt.Println("\nMemory:")
	fmt.Printf("   Total: %.1f GB\n", float64(memInfo.Total)/GB)
	fmt.Printf("   Used: %.1f GB\n", float64(memInfo.Used)/GB)
	fmt.Printf("   Used %%: %.1f%%\n", memInfo.UsedPercent)

	fmt.Println("\nDisk:")
	fmt.Printf("   Total: %.1f GB\n", float64(diskInfo.Total)/GB)
	fmt.Printf("   Used: %.1f GB\n", float64(diskInfo.Used)/GB)
	fmt.Printf("   Used %%: %.1f%%\n", diskInfo.UsedPercent)
}