package main

import (
	"log"
	"net/http"

	"github.com/go-backend-rest/internal/config"
	"github.com/go-backend-rest/internal/handler"
)

func main() {
	cfg := config.MustLoad()
	log.Printf("loaded config: %+v", cfg)
	handler := handler.Health
	mux := http.NewServeMux()

	mux.HandleFunc("/api/health", handler)
}
