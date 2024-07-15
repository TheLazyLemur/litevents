package main

import (
	"litevents/store"
	"log"
	"log/slog"
	"net/http"

	_ "embed"
)

var (
	port string = ":3000"
	s    store.Store
)

func main() {
	mux := http.NewServeMux()

	s = store.NewSqliteStore("litevents.db")

	server := http.Server{
		Addr:    port,
		Handler: mux,
	}

	slog.Info("Started http server", "port", port)
	log.Fatal(server.ListenAndServe())
}
