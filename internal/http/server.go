package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/message"
)

func PublishMessagesServer(publisher message.Publisher) {
	publishHandler := publishHandler{Publisher: publisher}
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/publish", publishHandler.publishHandler)
	fmt.Printf("Starting server for testing HTTP POST...\n")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
