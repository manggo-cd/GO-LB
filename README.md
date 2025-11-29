# Simple Load Balancer in Go

A simple yet functional HTTP load balancer written in Go that distributes incoming requests across multiple backend servers using round-robin algorithm.

## Features

- **Round-Robin Load Balancing**: Distributes requests evenly across backend servers
- **Health Checks**: Automatically monitors backend server health every 20 seconds
- **Automatic Failover**: Routes requests away from failed servers
- **Retry Mechanism**: Retries failed requests on other available servers
- **Reverse Proxy**: Uses Go's built-in `httputil.ReverseProxy` for efficient proxying

## Architecture

```
Client Request → Load Balancer (port 8080) → Backend Servers (8081, 8082, 8083)
```

The load balancer:
1. Receives incoming HTTP requests on port 8080
2. Selects the next healthy backend server using round-robin
3. Forwards the request to the selected backend
4. Returns the backend's response to the client
5. If a backend fails, automatically retries with another server

## Prerequisites

- Go 1.16 or higher installed on your system

## Installation

1. Clone or download this repository
2. Navigate to the project directory

## Usage

### Step 1: Start Backend Servers

Open three separate terminal windows and run:

**Terminal 1:**
```bash
wsl go run backend_server.go 8081
```

**Terminal 2:**
```bash
wsl go run backend_server.go 8082
```

**Terminal 3:**
```bash
wsl go run backend_server.go 8083
```

### Step 2: Start the Load Balancer

Open a fourth terminal window and run:

```bash
wsl go run main.go
```

You should see output like:
```
Configured backend: http://localhost:8081
Configured backend: http://localhost:8082
Configured backend: http://localhost:8083
Load Balancer started at :8080
```

### Step 3: Test the Load Balancer

Open a browser or use curl to send requests:

```bash
curl http://localhost:8080
```

Each request will be distributed to a different backend server in round-robin fashion.

To test multiple requests:

```bash
# In PowerShell
for ($i=1; $i -le 6; $i++) { curl http://localhost:8080 }
```

```bash
# In WSL/bash
for i in {1..6}; do curl http://localhost:8080; done
```

You should see responses from different backend servers (ports 8081, 8082, 8083) in rotation.

## Testing Failover

1. Start all three backend servers and the load balancer
2. Send some requests - you'll see all three backends responding
3. Stop one backend server (Ctrl+C in its terminal)
4. Continue sending requests - the load balancer will automatically stop routing to the failed server
5. Restart the stopped backend - after the next health check (20 seconds), it will be back in rotation

## Configuration

You can modify the backend servers in `main.go`:

```go
serverList := []string{
    "http://localhost:8081",
    "http://localhost:8082",
    "http://localhost:8083",
}
```

You can also adjust:
- **Load balancer port**: Change `Addr: ":8080"` in `main.go`
- **Health check interval**: Modify `time.NewTicker(20 * time.Second)` in the `healthCheck` function
- **Backend timeout**: Adjust `timeout := 2 * time.Second` in the `isBackendAlive` function
- **Retry attempts**: Change `retries := 3` in the error handler

## How It Works

### Round-Robin Algorithm
The load balancer maintains a counter that atomically increments with each request, ensuring even distribution across all healthy backends.

### Health Checks
Every 20 seconds, the load balancer sends HTTP requests to all backends and updates their alive status based on the response (status code < 500 means alive).

### Error Handling
If a request to a backend fails:
1. The load balancer marks it as an error
2. Attempts to retry with up to 3 different backends
3. Returns a 503 error if no backends are available

## Project Structure

```
.
├── main.go              # Load balancer implementation
├── backend_server.go    # Simple backend server for testing
├── go.mod              # Go module file
└── README.md           # This file
```

## Future Enhancements

Possible improvements to this load balancer:
- Add weighted round-robin
- Implement least-connections algorithm
- Add SSL/TLS support
- Implement sticky sessions
- Add request/response logging
- Add metrics and monitoring dashboard
- Support for dynamic backend addition/removal
- Configuration file support (JSON/YAML)

## License

This is a simple educational project. Feel free to use and modify as needed.

