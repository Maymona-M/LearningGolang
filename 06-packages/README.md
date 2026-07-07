# 06-packages - Go Packages Practice

This section introduces how Go organizes and reuses code through packages.

Packages are a core part of Go development because they allow programs to be separated into smaller, maintainable modules. Instead of keeping everything inside one `main.go` file, functionality can be divided into different folders and imported when needed.

## Topics Covered

* Understanding Go packages
* Standard library packages
* Importing packages
* Creating and using external packages
* Managing dependencies with Go modules
* Using `go get`
* Using `go mod tidy`
* Reading package documentation
* Accessing system information through third-party libraries

## Folder Structure

```
06-packages/

├── fmt/
│   └── main.go
│
├── time/
│   └── main.go
│
├── os/
│   ├── main.go
│   └── example.txt
│
├── bufio/
│   └── main.go
│
└── third-party/
    └── main.go
```

## Packages Practiced

### fmt

Used for formatted output.

Examples:

* Printing values
* Formatting strings
* Using `Println()` and `Printf()`

---

### time

Used for working with time-related operations.

Examples:

* Delays using `time.Sleep()`
* Retrieving current timestamps
* Creating timed processes

This will be useful for the system monitor project to refresh statistics periodically.

---

### os

Used for interacting with the operating system.

Examples:

* Opening files
* Working with system resources
* Understanding functions that return errors

Example:

```go
file, err := os.Open("example.txt")
```

This introduced Go's common error handling pattern:

```go
if err != nil {
    return
}
```

---

### bufio

Used for buffered input and reading data efficiently.

Examples:

* Reading user input
* Creating readers using `bufio.NewReader()`

---

### Third-party packages

External packages provide functionality that is not included in Go's standard library.

Installed and used:

```
github.com/shirou/gopsutil/v3
```

The `gopsutil` package allows Go programs to access operating system statistics.

Used features:

* Memory monitoring
* CPU statistics
* Disk information
* Process information

Example:

```go
memory, err := mem.VirtualMemory()

fmt.Println(memory.UsedPercent)
```

## Go Module Commands

### Initialize a Go module

```bash
go mod init learninggolang
```

Creates:

```
go.mod
```

which tracks the project module and dependencies.

---

### Install external packages

Example:

```bash
go get github.com/shirou/gopsutil/v3
```

Downloads and adds third-party dependencies.

---

### Clean and manage dependencies

```bash
go mod tidy
```

Used to:

* Add missing dependencies
* Remove unused dependencies
* Update `go.mod` and `go.sum`

---

## Key Concepts Learned

### Importing packages

Packages are imported using:

```go
import "package-name"
```

Example:

```go
import (
    "fmt"
    "time"
)
```

---

### Package organization

Go projects are usually structured by responsibility:

```
project/

main.go

cpu/
    cpu.go

memory/
    memory.go

disk/
    disk.go
```

Each package handles one specific task.

---


## Notes

* Standard library packages are included with Go.
* Third-party packages are downloaded using `go get`.
* Dependencies should be managed using `go mod tidy`.
* Keep `go.mod` and `go.sum` in version control.
* Avoid spaces in folder/package names because Go uses folder paths as package identifiers.

Each example can be run using:

```bash
go run .
```
