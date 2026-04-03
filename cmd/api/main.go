package main

import (
	"log"
	"net/http"
	"os"

	"github.com/calculator-api-go-en/calculator-api/internal/server"
	"github.com/calculator-api-go-en/calculator-api/internal/usecases"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	calculator := usecases.NewCalculator()
	handler := server.NewRouter(calculator)

	addr := ":" + port
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
