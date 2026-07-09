package main

import (
	"fmt"
	"log"
	"net"

	"learninggolang/projects/1_system_monitor/cpu"
	"learninggolang/projects/1_system_monitor/disk"
	"learninggolang/projects/1_system_monitor/memory"
)


func main() {

	cpuInfo, err := cpu.GetCPU()
	if err != nil {
		fmt.Println("Error retrieving CPU info: ", err)
		return
	}

	memInfo, err := memory.GetMemory()
	if err != nil {
		fmt.Println("Error retrieving memory info: ", err)
		return
	}

	diskInfo, err := disk.GetDisk("/")
	if err != nil {
		fmt.Println("Error retrieving disk info: ", err)
		return
	}

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Connected to server!")

	
	msg := fmt.Sprintf(
		"CPU: %.1f%% | Memory: %.1f%% | Disk: %.1f%%\n",
		cpuInfo.Usage,
		memInfo.UsedPercent,
		diskInfo.UsedPercent,
	)

	fmt.Fprint(conn, msg)
}