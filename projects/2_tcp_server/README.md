# 2: TCP Server

A finished mini-project: a TCP-based system agent that collects live CPU, Memory, and Disk usage and sends it to a server over a network connection — with the agent running as a background service on Linux, and readings persisted to a SQLite database on the server.

## Project Structure

```
2_tcp_server/
├── agent.go     # Collects system stats and sends them to the server
├── server.go     # Listens for connections, decodes stats, logs + stores them
├── monitor.db     # SQLite database (gitignored, created at runtime)
└── stats/
    └── stats.go     # Shared SystemStats struct used by both agent and server
```

## How It Works

**`stats/stats.go`**
- Defines a shared `SystemStats` struct (`CPU`, `Mem`, `Disk`, `Time`) so both agent and server encode/decode the exact same shape

**`server.go`**
1. Opens (or creates) `monitor.db` and ensures a `readings` table exists
2. Starts a TCP listener on port `8080`
3. For every incoming connection, spins up a goroutine (`go handleConnection(conn)`) so multiple agents can send data at the same time
4. Reads the sender's IP via `conn.RemoteAddr()`
5. Decodes the incoming `gob`-encoded `SystemStats` struct
6. Prints a readable log line and inserts the reading (IP, CPU, Mem, Disk, Unix timestamp) into SQLite
7. Sends back an ACK

**`agent.go`**
1. Collects live stats using the `1_system_monitor` packages (`cpu`, `memory`, `disk`)
2. Packs the values into a `SystemStats` struct
3. Connects to the server's real IP (not `localhost`) on port `8080`
4. Encodes the struct with `encoding/gob` and sends it over the connection
5. Half-closes its write side (`CloseWrite()`), waits for the ACK, closes the connection
6. Sleeps 3 seconds and repeats — runs continuously

## Key Concepts Used

| Concept | Meaning |
|---|---|
| `struct` | Groups CPU, Mem, Disk, and Time into one typed object instead of loose variables |
| `encoding/gob` | Go's native binary format — encodes/decodes the struct directly to bytes; matches fields by **name**, not position, so both sides need identical field names/types |
| `net.Listen` / `net.Dial` | Server-side listening vs. client-side connecting |
| `go handleConnection(conn)` | Goroutine — lets the server handle multiple simultaneous client connections concurrently |
| `conn.RemoteAddr()` | Identifies the IP (and port) of whoever connected; used to tag each stored reading with its source |
| `database/sql` + `modernc.org/sqlite` | Pure-Go SQLite driver (no C compiler needed) used to persist readings to a local `.db` file |
| systemd unit file | Config that runs the agent as a managed Linux background service — starts on boot, restarts automatically on crash |

## How to Run

**Server (Windows or wherever you want data collected):**
```bash
cd projects/2_tcp_server
go run server.go
```

**Agent (edit the dial IP to point at the server's real address first):**
```bash
cd projects/2_tcp_server
go build -o agent agent.go
./agent
```

Expected server output:
```
Server started...
Waiting for connections...
Client connected from: 192.168.56.102
From 192.168.56.102 -> CPU: 2.5% | Mem: 17.9% | Disk: 32.5% | Time: 1784447200
```

**Running the agent as a systemd service (Linux):**
```bash
sudo cp agent /usr/local/bin/agent
sudo nano /etc/systemd/system/agent.service
```
```ini
[Unit]
Description=Go System Monitor Agent
After=network.target

[Service]
ExecStart=/usr/local/bin/agent
Restart=always
RestartSec=3
User=<your-username>

[Install]
WantedBy=multi-user.target
```
```bash
sudo systemctl daemon-reload
sudo systemctl enable agent.service
sudo systemctl start agent.service
```

**Viewing stored data:**
Open `monitor.db` with [DB Browser for SQLite](https://sqlitebrowser.org/), or query:
```sql
SELECT id, ip, cpu, mem, disk, datetime(timestamp, 'unixepoch') AS readable_time
FROM readings ORDER BY id DESC;
```

## Notes

- Import paths rely on the root module `learninggolang`, importing `1_system_monitor`'s `cpu`, `memory`, and `disk` packages directly — no duplicated logic
- The agent always dials the **real IP** of whichever machine runs the server — never `localhost` — so agent and server can run on physically separate machines
- VirtualBox **NAT** isolates a VM from the host; **Bridged** (same physical LAN) or **Host-Only** (private virtual network, stable IP regardless of Wi-Fi changes) is needed for direct communication
- `log.Fatal(err)` on a failed `Dial` kills the whole agent process — systemd's `Restart=always` recovers it, but retrying the loop internally (without exiting) would be a cleaner design
- Mobile hotspots often enable client/AP isolation, silently blocking device-to-device traffic even on the "same" Wi-Fi — a common cause of connection timeouts that isn't a code or firewall issue
```