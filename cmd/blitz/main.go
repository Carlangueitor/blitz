package main

import (
	"log"

	"github.com/carlangueitor/blitz"
	"github.com/carlangueitor/blitz/configloader"
	"github.com/carlangueitor/blitz/websocket"
)

func start(configLoader blitz.ConfigLoader, server blitz.Server) {
	config, err := configLoader.Load()
	if err != nil {
		log.Printf("Error loading config: %s\n", err)
	}

	log.Printf("Config Loaded: %+v\n", config)

	server.SetConfig(config)
	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	viperConfigLoader := configloader.ViperConfigLoader{}
	webSocketServer := websocket.Server{}
	start(&viperConfigLoader, &webSocketServer)
}
