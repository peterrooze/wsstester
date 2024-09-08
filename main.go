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
	numConnections := flag.Int("c", 1, "Number of connections to make")
	flag.Parse()

	if *serverURL == "" {
		log.Fatal("Please provide a WebSocket server URL")
	}

	u, err := url.Parse(*serverURL)
	if err != nil {
		log.Fatal("Invalid URL:", err)
	}

	var (
		successCount uint32
		failCount    uint32
		wg           sync.WaitGroup
	)

	fmt.Println("Dialing: ", u.String())
	for i := 0; i < *numConnections; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				atomic.AddUint32(&failCount, 1)
				return
			}
			defer c.Close()
			atomic.AddUint32(&successCount, 1)
		}()
	}

	wg.Wait()

	fmt.Printf("Results:\n")
	fmt.Printf("Successful connections: %d\n", successCount)
	fmt.Printf("Failed connections: %d\n", failCount)
}
