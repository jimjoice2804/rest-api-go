package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-backend-rest/internal/config"
	"github.com/go-backend-rest/internal/handler"
)

func main() {
	cfg := config.MustLoad()
	// log.Printf("loaded config: %+v", cfg)
	addr := cfg.HTTP.Address
	handler := handler.Health
	mux := http.NewServeMux()

	mux.HandleFunc("/api/health", handler)

	//fallback
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	//start a server
	fmt.Println(addr)
	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("Something went wrong", err)
	}
}
