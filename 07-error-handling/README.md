# 07-error-handling - Go Error Handling Practice

This section introduces how Go handles errors and why error handling is a core part of writing reliable Go programs.

Unlike languages that use exceptions (`try/catch`), Go treats errors as normal return values. Functions commonly return two values: the result and an error. The programmer checks the error immediately and decides how to handle it.

## Topics Covered

* Understanding the `error` type
* Handling errors returned from Go packages
* Using `if err != nil`
* Returning errors from custom functions
* Creating errors using the `errors` package
* Understanding `nil` errors
* Writing safer and more reliable functions

## Folder Structure

```text
07-error-handling/

├── basic_error.go
├── error_return.go
└── README.md
```

## Concepts Practiced

### Basic Error Handling

Go functions often return a value and an error together.

Example:

```go
file, err := os.Open("example.txt")

if err != nil {
    fmt.Println(err)
    return
}
```

The returned values represent:

* The successful result
* Any error that occurred

If `err` is `nil`, the operation was successful.

---

### Returning Errors From Functions

Functions can return their own errors:

```go
func divide(a int, b int) (int, error)
```

Example:

```go
if b == 0 {
    return 0, errors.New("cannot divide by zero")
}
```

This allows functions to communicate problems back to the caller.

---

## Common Go Error Pattern

A large amount of Go code follows this structure:

```go
result, err := someFunction()

if err != nil {
    // handle error
    return
}
```

Examples:

```go
file, err := os.Open()

memory, err := mem.VirtualMemory()
```

---

## The `errors` Package

Go provides the built-in `errors` package to create custom errors.

Import:

```go
import "errors"
```

Create an error:

```go
errors.New("something went wrong")
```

---

## Why Error Handling Matters

Error handling prevents programs from silently failing.

Examples of possible errors:

* File does not exist
* Invalid user input
* Network failure
* Permission issues
* System resource unavailable

Instead of crashing unexpectedly, Go programs can detect problems and respond appropriately.

---


## Notes

* Go does not use traditional `try/catch` for errors.
* Errors are returned as values.
* Always check errors after calling functions that can fail.
* Returning errors makes code easier to debug and maintain.

Each example can be run using:

```bash
go run .
```
