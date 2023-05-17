package websocket

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func newServer(handler http.Handler) (*httptest.Server, string) {
	server := httptest.NewServer(handler)
	url := "ws" + server.URL[4:]
	return server, url
}

func getConnection(t *testing.T, handler http.Handler) (*httptest.Server, net.Conn) {
	server, url := newServer(handler)

	conn, _, _, err := ws.Dial(context.Background(), url)
	if err != nil {
		t.Fatalf("Failed to dial WebSocket: %v", err)
	}

	return server, conn
}

func TestWebSocketHandlerServeHTTPUpgradeConnection(t *testing.T) {
	handler := webSocketHandler{}
	server, url := newServer(handler)
	defer server.Close()

	conn, _, _, err := ws.Dial(context.Background(), url)
	if err != nil {
		t.Fatalf("Failed to dial WebSocket: %v", err)
	}
	defer conn.Close()
}

func TestWebSocketHandlerGetSuscriberID(t *testing.T) {
	server, conn := getConnection(t, NewWebSocketHandler())
	defer server.Close()
	defer conn.Close()

	msg, op, err := wsutil.ReadServerData(conn)
	if err != nil {
		t.Fatalf("Failed to read server data from WebSocket: %v", err)
	}

	if op != ws.OpText {
		t.Errorf("Unexpected message type. Expected: %v, Got: %v", ws.OpText, op)
	}

	response := message{}
	err = json.Unmarshal(msg, &response)
	if err != nil {
		t.Fatalf("Received non valid JSON: %v", err)
	}

	if response.Op != OpSubscribeId {
		t.Errorf("Unexpected message Op. Expected: %s, Got: %s", OpSubscribeId, response.Op)
	}
}

func TestWebSocketHandlerJSONParsing(t *testing.T) {
	handler := webSocketHandler{}
	server, conn := getConnection(t, handler)
	defer server.Close()
	defer conn.Close()

}
