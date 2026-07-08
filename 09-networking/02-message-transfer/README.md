# 02-message-transfer

Building on `01-tcp-basics` — now the client and server actually exchange data instead of just connecting.

## Goal

Learn how to send and receive real data over an established TCP connection.

## Files

```
02-message-transfer/
├── server.go
└── client.go
```

## How It Works

**server.go**
1. Does everything from `01-tcp-basics` (listen, accept)
2. **Reads** a message sent by the client
3. Prints the received message

**client.go**
1. Does everything from `01-tcp-basics` (connect)
2. **Writes** a message to the server

## Key Concepts

| Concept | Meaning |
|---|---|
| `bufio.NewReader(conn)` | Wraps the connection so you can read text from it easily |
| `.ReadString('\n')` | Reads everything up to (and including) the first newline — i.e. "read one line" |
| `fmt.Fprintln(conn, "message")` | Like `fmt.Println`, but writes to the connection instead of the terminal; the `ln` adds a newline automatically |

The newline (`\n`) matters here — the server reads *until* it sees one, so the client must send one for the message to arrive correctly.

## How to Run

**Terminal 1 — server:**
```bash
go run server.go
```
```
Server started...
Waiting for connection...
```

**Terminal 2 — client:**
```bash
go run client.go
```
```
Connected to server!
```

**Back in Terminal 1:**
```
Client connected!
Message received: Hello Server
```

## Milestone

- Client sends real data over TCP
- Server reads and displays it
- First hands-on use of `Read`/`Write` over a connection
