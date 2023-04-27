package main

import (
	"fmt"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"

	"github.com/carlangueitor/blitz"
	"github.com/carlangueitor/blitz/configloader"
)

func start(configLoader blitz.ConfigLoader) {
	config, err := configLoader.Load()
	if err != nil {
		fmt.Printf("Error loading config: %s", err)
	}

	fmt.Printf("Port %s", config.Port)

	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Serving at %s", config.Port)
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// handle error
		}
		go func() {
			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					// handle error
				}
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					// handle error
				}
			}
		}()
	}))
}

func main() {
	viperConfigLoader := configloader.ViperConfigLoader{}
	start(&viperConfigLoader)
}
