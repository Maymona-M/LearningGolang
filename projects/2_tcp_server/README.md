# 2_tcp_server

A finished mini-project: a TCP-based system agent that collects live CPU, Memory, and Disk usage and sends it to a server over a network connection.

## 📁 Project Structure

```
2_tcp_server/
├── agent/
│   └── agent.go     # Collects system stats and sends them to the server
└── server/
    └── server.go     # Listens for a connection and prints received stats
```

## ⚙️ How It Works

**`server/server.go`**
1. Starts a TCP listener on port `8080`
2. Waits for a connection
3. Reads the incoming message
4. Prints it to the terminal

**`agent/agent.go`**
1. Collects live stats using the `1_system_monitor` packages (`cpu`, `memory`, `disk`)
2. Connects to the server at `localhost:8080`
3. Formats the stats as a single line of text
4. Sends it over the connection

## 🧠 Key Concepts Used

| Concept | Meaning |
|---|---|
| `net.Listen` / `net.Dial` | Server-side listening vs. client-side connecting |
| `bufio.NewReader(conn).ReadString('\n')` | Reads one line of text from a connection |
| `fmt.Fprintf(conn, ...)` | Formats text and writes it directly to a connection |
| `log.Fatal(err)` | Used for setup failures where the program truly cannot continue |

## 🚀 How to Run

**Terminal 1 — start the server:**
```bash
cd projects/2_tcp_server/server
go run server.go
```

**Terminal 2 — start the agent:**
```bash
cd projects/2_tcp_server/agent
go run agent.go
```

Expected server output:
```
Server started...
Waiting for connection...
Client connected!
Message received: CPU: 45.2% | Memory: 77.7% | Disk: 62.3%
```

## 🔧 Notes

- This project currently handles **one connection at a time** — the server exits after receiving one message, and the agent sends only once
- Import paths rely on the root module `learninggolang`, importing `1_system_monitor`'s `cpu`, `memory`, and `disk` packages directly — no duplicated logic