# 01-tcp-basics

The first step into networking with Go — establishing a raw TCP connection between two separate programs.

## Goal

Understand the absolute basics of TCP communication:
- What is a TCP server?
- What is a TCP client?
- How does a connection get established?

No data is sent yet — this step only proves that a server and client can find each other and connect.

## Files

```
01-tcp-basics/
├── server.go
└── client.go
```

## How It Works

**server.go**
1. Creates a TCP listener on port `8080`
2. Waits for a client to connect (`Accept()` blocks until someone connects)
3. Prints `"Client connected!"`

**client.go**
1. Connects to `localhost:8080`
2. Prints `"Connected to server!"`

## Key Concepts

| Concept | Meaning |
|---|---|
| `net.Listen("tcp", ":8080")` | Opens port `8080` and starts listening for TCP connections |
| `listener.Accept()` | Blocks (pauses) until a client connects |
| `net.Dial("tcp", "localhost:8080")` | Client-side: actively connects to a running server |
| `defer conn.Close()` | Ensures the connection is closed when the function ends, no matter what |
| `log.Fatal(err)` | Prints the error and exits immediately — used here because if the server can't even start listening, there's nothing left to do |

`localhost:8080` breaks down as:
- `localhost` → your own computer
- `8080` → the "door number" (port) the server is listening on

## How to Run

Two separate terminals are needed — server and client run as independent programs.

**Terminal 1 — start the server:**
```bash
cd 09-networking/01-tcp-basics
go run server.go
```

Output:
```
Server started...
Waiting for connection...
```

**Terminal 2 — start the client:**
```bash
cd 09-networking/01-tcp-basics
go run client.go
```

Output:
```
Connected to server!
```

**Back in Terminal 1**, it immediately unblocks and prints:
```
Client connected!
```

## Milestone

- Two independent Go programs
- Running in separate terminals
- Successfully talking to each other over TCP

