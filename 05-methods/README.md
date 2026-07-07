# 05-methods - Go Methods Practice

This folder covers how to attach behaviour to structs using methods in Go.

## Topics Covered

- Basic methods (value receivers)
- Pointer receivers and mutating structs
- The `String()` method and custom print formatting

## Purpose

This section introduces object-like behaviour in Go. It teaches how a struct can own the logic that operates on it, rather than relying only on free-standing functions, and clarifies when a method should modify the original struct versus a copy. These concepts are a stepping stone toward interfaces and are directly useful for the system monitor project, where structs like `Stats` will need methods for formatting and displaying collected data.

## Notes

- All files use `package main` for simplicity.
- Each file is independent and can be run using:

```bash
go run filename.go
```