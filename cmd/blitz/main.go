package main

import (
	"fmt"
	"log"

	"github.com/carlangueitor/blitz"
	"github.com/carlangueitor/blitz/configloader"
	"github.com/carlangueitor/blitz/websocket"
)

func start(configLoader blitz.ConfigLoader, server blitz.Server) {
	config, err := configLoader.Load()
	if err != nil {
		fmt.Printf("Error loading config: %s", err)
	}

	fmt.Printf("Config Loaded: %+v\n", config)

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
