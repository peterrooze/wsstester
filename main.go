package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"sync"
	"sync/atomic"

	"github.com/gorilla/websocket"
)

func main() {
	serverURL := flag.String("u", "", "WebSocket server URL (wss://)")
	numConnections := flag.Int("c", 1, "Number of connections per client")
	numClients := flag.Int("n", 1, "Number of clients")
	flag.Parse()

	if *serverURL == "" {
		log.Fatal("Please provide a WebSocket server URL")
	}

	success, fail := runConnectionTest(*serverURL, *numClients, *numConnections)

	fmt.Printf("Results:\n")
	fmt.Printf("Successful connections: %d\n", success)
	fmt.Printf("Failed connections: %d\n", fail)
}

func runConnectionTest(serverURL string, numClients, numConnections int) (uint32, uint32) {
	u, err := url.Parse(serverURL)
	if err != nil {
		log.Fatal("Invalid URL:", err)
	}

	var (
		successCount uint32
		failCount    uint32
		wg           sync.WaitGroup
	)

	fmt.Printf("Dialing: %s with %d clients, each making %d connections\n", u.String(), numClients, numConnections)

	for client := 0; client < numClients; client++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()
			for i := 0; i < numConnections; i++ {
				c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
				if err != nil {
					atomic.AddUint32(&failCount, 1)
					continue
				}
				defer c.Close()
				atomic.AddUint32(&successCount, 1)
			}
		}(client)
	}

	wg.Wait()

	return successCount, failCount
}
