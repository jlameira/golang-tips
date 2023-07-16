package http

import (
	"log"
	"net/http"

	"cmd/api/main.go/internal/handlers/frete"
	infrastructure "cmd/api/main.go/internal/infraestruture"
	"github.com/gorilla/mux"

	"cmd/api/main.go/internal/application"
)

// Server representa um servidor HTTP
type Server struct {
	calculadoraService *application.CalculaFreteService
}

// NewServer cria uma instância do servidor HTTP
func NewServer(calculadoraService *application.CalculaFreteService) *Server {
	return &Server{
		calculadoraService: calculadoraService,
	}
}

// StartServer inicia o servidor HTTP na porta especificada
func StartServer() {
	router := mux.NewRouter()

	gateway := infrastructure.NewCorreiosAPI()
	calculadoraService := application.NewCalculaFreteService(gateway)

	server := NewServer(calculadoraService)
	server.registerRoutes(router, calculadoraService)

	addr := ":8080"
	log.Printf("Servidor iniciado em %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func (s *Server) registerRoutes(router *mux.Router, service *application.CalculaFreteService) {
	router.HandleFunc("/calcula-frete", frete.NewHandlerFrete(service).CalculateFreightHandler).Methods("POST")
}
