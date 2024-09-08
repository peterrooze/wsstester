# WebSocket Connection Tester

This Go program allows you to test multiple WebSocket connections from multiple clients to a specified server simultaneously. It's useful for benchmarking WebSocket server performance or testing connection limits.

## Features

- Establish multiple WebSocket connections concurrently
- Configurable number of connections per client
- Configurable number of clients
- Reports successful and failed connection attempts

## Usage

```
./wsstester -u <WebSocket_URL> -c <number_of_connections> -n <number_of_clients>

```

## Build from Source prerequisites

- Go 1.15 or higher
- `github.com/gorilla/websocket` package

## Installation

1. Clone this repository:
   ```
   git clone https://github.com/peterrooze/wsstester.git
   cd wsstester
   ```

2. Install the required dependency:
   ```
   go get github.com/gorilla/websocket
   ```

## Usage

Run the program with the following command:

```
go run main.go -u <WebSocket_URL> -c <number_of_connections> -n <number_of_clients>
```

### Flags

- `-u`: WebSocket server URL (required, must start with `wss://` or `ws://`)
- `-c`: Number of connections to make (default: 1)
- `-n`: Number of clients

### Example

```
build/wsstester -u "wss://demo.piesocket.com/v3/channel_123?api_key=VCXCEuvhGcBDP7XhiJJUDvR1e1D3eiVjgZ9VRiaV&notify_self" -c 10 -n 5
```