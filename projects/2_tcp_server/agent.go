package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"time"

	"learninggolang/projects/1_system_monitor/cpu"
	"learninggolang/projects/1_system_monitor/disk"
	"learninggolang/projects/1_system_monitor/memory"

	"learninggolang/projects/2_tcp_server/stats"
)

func main() {
	for {
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

		payload := stats.SystemStats{
			CPU:  cpuInfo.Usage,
			Mem:  memInfo.UsedPercent,
			Disk: diskInfo.UsedPercent,
			Time: time.Now().Unix(),
		}

		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connected to server!")

		encoder := gob.NewEncoder(conn)
		if err := encoder.Encode(payload); err != nil {
			fmt.Println("Error encoding stats:", err)
		}

		tcpConn := conn.(*net.TCPConn)
		tcpConn.CloseWrite()

		ack := make([]byte, 128)
		n, err := conn.Read(ack)
		if err != nil {
			fmt.Println("Error reading ACK:", err)
		} else {
			fmt.Println("Server says:", string(ack[:n]))
		}

		conn.Close()
		time.Sleep(3 * time.Second)
	}
}
