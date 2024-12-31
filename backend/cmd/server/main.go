package main

import (
	"log"
	"net/http"
	"rfid_handler/pkg/handler"
	"rfid_handler/pkg/state"
)

func main() {
    appState := state.NewAppState()

    http.HandleFunc("/rfid/collect", handler.HandleUpdateCollected(appState))

    log.Println("Server starting on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
