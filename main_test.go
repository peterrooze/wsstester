package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func TestRunConnectionTest(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer c.Close()
		// Echo back any messages received
		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				break
			}
		}
	}))
	defer server.Close()

	// Replace "http" with "ws" in the test server URL
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	tests := []struct {
		name           string
		numClients     int
		numConnections int
		wantSuccess    uint32
		wantFail       uint32
	}{
		{"Single client, single connection", 1, 1, 1, 0},
		{"Single client, multiple connections", 1, 5, 5, 0},
		{"Multiple clients, single connection", 3, 1, 3, 0},
		{"Multiple clients, multiple connections", 3, 3, 9, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSuccess, gotFail := runConnectionTest(wsURL, tt.numClients, tt.numConnections)
			if gotSuccess != tt.wantSuccess {
				t.Errorf("runConnectionTest() gotSuccess = %v, want %v", gotSuccess, tt.wantSuccess)
			}
			if gotFail != tt.wantFail {
				t.Errorf("runConnectionTest() gotFail = %v, want %v", gotFail, tt.wantFail)
			}
		})
	}
}
