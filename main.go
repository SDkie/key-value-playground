package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SDkie/key-value-playground/config"
	"github.com/SDkie/key-value-playground/db"
	"github.com/SDkie/key-value-playground/handler"
)

func main() {
	cfg := config.Init()

	err := db.Init()
	if err != nil {
		log.Panic(err)
	}

	http.HandleFunc("/keyvalue", handler.KeyValue)

	addr := fmt.Sprintf("localhost:%s", cfg.Webserver.WEBSERVER_PORT)
	log.Println("Starting server at -", addr)
	log.Printf("KeyValue WebSocket running at - ws://%s/keyvalue\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
