package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/carlangueitor/blitz"
)

type Server struct {
	config *blitz.Config
}

func (server *Server) SetConfig(config *blitz.Config) {
	server.config = config
}

func (server Server) Start() error {
	mux := http.NewServeMux()
	mux.Handle("/", NewWebSocketHandler())

	log.Printf("Serve config %v\n", server)

	listenAddr := fmt.Sprintf(":%d", server.config.Port)
	err := http.ListenAndServe(listenAddr, mux)
	return err
}
