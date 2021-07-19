package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

type publishHandler struct {
	Publisher message.Publisher
}

type Payload struct {
	LogTime  string      `json:"log_time"`
	Level    string      `json:"level"`
	Name     string      `json:"name"`
	Message  interface{} `json:"message"`
	Hostname string      `json:"hostname"`
	Tag      string      `json:"tag"`
}

func (p publishHandler) publishHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/publish" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var payload Payload
		err := decoder.Decode(&payload)
		if err != nil {
			panic(err)
		}
		msgdata, _ := json.Marshal(payload.Message)
		msg := message.NewMessage(watermill.NewUUID(), []byte(msgdata))
		middleware.SetCorrelationID(watermill.NewUUID(), msg)
		fmt.Printf("%+v", msg)
		fmt.Printf("\n\n\nSending message %s, correlation id: %s\n", msg.UUID, middleware.MessageCorrelationID(msg))
		if err := p.Publisher.Publish(payload.Tag, msg); err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(struct{ Status string }{Status: "Success"})
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ping" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(struct{ Status string }{Status: "Success"})
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
