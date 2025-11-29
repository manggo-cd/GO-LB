# Simple Load Balancer in Go

A simple yet functional HTTP load balancer written in Go that distributes incoming requests across multiple backend servers using round-robin algorithm.

## Features

- **Round-Robin Load Balancing**: Distributes requests evenly across backend servers
- **Health Checks**: Automatically monitors backend server health every 20 seconds
- **Automatic Failover**: Routes requests away from failed servers
- **Retry Mechanism**: Retries failed requests on other available servers
- **Reverse Proxy**: Uses Go's built-in `httputil.ReverseProxy` for efficient proxying
