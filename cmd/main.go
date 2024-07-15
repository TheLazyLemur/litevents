package main

import (
	"encoding/json"
	"fmt"
	"litevents/store"
	"litevents/types"
	"log"
	"log/slog"
	"net/http"

	_ "embed"

	"github.com/gorilla/websocket"
)

var (
	port string = ":3000"
	s    store.Store

	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func main() {
	s = store.NewSqliteStore("litevents.db")

	mux := http.NewServeMux()

	mux.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			slog.Error("Failed to upgrade to websocket", "err", err)
			return
		}

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				slog.Error("Failed to read message", "err", err)
				return
			}

			var cMsg types.Message
			err = json.Unmarshal(msg, &cMsg)
			if err != nil {
				slog.Error("Failed to unmarshal message", "err", err)
				break
			}

			switch cMsg.Type {
			case "consume":
				fmt.Println("Consume")
			case "ack":
				fmt.Println("Ack")
			default:
				slog.Error("Unknown message type", "type", cMsg.Type)
			}
		}
	})

	server := http.Server{
		Addr:    port,
		Handler: mux,
	}

	slog.Info("Started http server", "port", port)
	log.Fatal(server.ListenAndServe())
}
