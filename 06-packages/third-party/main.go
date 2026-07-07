package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
)

func main(){

	memory, err := mem.VirtualMemory()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("RAM Usage: %.2f%%\n", memory.UsedPercent)

	fmt.Printf(
		"Total RAM: %.2f GB\n",
		float64(memory.Total)/1024/1024/1024,
	)

	fmt.Printf(
		"Used RAM: %.2f GB\n",
		float64(memory.Used)/1024/1024/1024,
	)

	fmt.Printf(
		"Available RAM: %.2f GB\n",
		float64(memory.Available)/1024/1024/1024,
	)

	fmt.Printf(
		"Free RAM: %.2f GB\n",
		float64(memory.Free)/1024/1024/1024,
	)

}