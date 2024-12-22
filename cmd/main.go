package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nais2008/project_go_yandex/internal/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/api/v1/calculate", handlers.CalculateHandler)

	log.Printf("Server is running on http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
