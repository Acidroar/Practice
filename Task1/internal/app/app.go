package app

import (
	"github.com/kenedyCO/Practice/internal/clients"
	"github.com/kenedyCO/Practice/internal/services"
	"github.com/kenedyCO/Practice/internal/transport"
)

const httpPort = ":8080"

func Run() {
	// Start clients
	client := clients.New()
	// Start services
	service := services.New(client)
	// Start http server
	httpServer := transport.New(httpPort, service)

	httpServer.AddRoute()
	httpServer.StartHttpServer()
}
