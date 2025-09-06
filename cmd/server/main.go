package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Fryuni/world-anvil-cli/pkg/config"
)

func main() {
	cfg := config.Load()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("World Anvil Server"))
	})

	port := cfg.GetPort()
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
