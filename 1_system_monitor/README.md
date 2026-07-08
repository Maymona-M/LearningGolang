# System Monitor

A simple Go CLI tool that reads your computer's **CPU**, **Memory**, and **Disk** usage and prints a clean summary to the terminal.

## Project Structure

```
system-monitor/
├── main.go              # Entry point - calls all three packages and prints results
├── cpu/
│   └── cpu.go            # Reads CPU usage %, core count, and model name
├── memory/
│   └── memory.go         # Reads total/used RAM and usage %
└── disk/
    └── disk.go            # Reads total/used disk space and usage %
```

## How It Works

Each folder (`cpu`, `memory`, `disk`) is its own **package** — a self-contained toolbox responsible for fetching one type of system data using the [gopsutil](https://github.com/shirou/gopsutil) library under the hood.

`main.go` imports all three packages, calls their `Get...()` functions, and prints the results in a readable bullet-point format.

## Requirements

- Go 1.20+ installed
- Internet access (only needed once, to download the gopsutil dependency)

## Setup & Run

```bash
# 1. Initialize the Go module (only needed once)
go mod init system-monitor

# 2. Install the gopsutil dependency
go get github.com/shirou/gopsutil/v3

# 3. Tidy up go.mod/go.sum
go mod tidy

# 4. Run the program
go run main.go
```

## Example Output

```
===== System Information =====

CPU:
  * Usage: 100.0%
  * Cores: 8
  * Model: Intel(R) Core(TM) i7-8650U CPU @ 1.90GHz

Memory:
  * Total: 23.8 GB
  * Used: 18.5 GB
  * Used %: 77.0%

Disk:
  * Total: 930.0 GB
  * Used: 579.0 GB
  * Used %: 62.3%
```
