package websocket

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func TestWebSocketHandler_ServeHTTP(t *testing.T) {
	handler := webSocketHandler{}

	// Create a test server with the WebSocket handler
	server := httptest.NewServer(handler)
	defer server.Close()

	// Create a WebSocket connection to the test server
	url := "ws" + server.URL[4:]
	conn, _, _, err := ws.Dial(context.Background(), url)
	if err != nil {
		t.Fatalf("Failed to dial WebSocket: %v", err)
	}
	defer conn.Close()

	// Write a message to the WebSocket connection
	message := []byte("Hello, WebSocket!")
	err = wsutil.WriteClientText(conn, message)
	if err != nil {
		t.Fatalf("Failed to write message to WebSocket: %v", err)
	}

	// Read the server's response from the WebSocket connection
	msg, op, err := wsutil.ReadServerData(conn)
	if err != nil {
		t.Fatalf("Failed to read server data from WebSocket: %v", err)
	}

	if op != ws.OpText {
		t.Errorf("Unexpected message type. Expected: %v, Got: %v", ws.OpText, op)
	}

	response := msg
	if string(response) != string(message) {
		t.Errorf("Unexpected message. Expected: %s, Got: %s", message, response)
	}
}
