package main

import (
	"log"

	"cmd/api/main.go/internal/adapters/http"
)

func main() {
	http.StartServer()
	log.Println("Servidor HTTP iniciado.")
}
