package transport

import (
	"log"

	"github.com/labstack/echo/v4"
)

type HttpServer struct {
	address string
	*echo.Echo
	service Service
}

func New(address string, service Service) *HttpServer {
	echoServer := echo.New()

	return &HttpServer{
		address: address,
		Echo:    echoServer,
		service: service,
	}
}

func (h *HttpServer) StartHttpServer() {
	err := h.Start(h.address)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *HttpServer) AddRoute() {
	version := h.Group("/v1")
	{
		version.GET("/", h.randFilm)
		version.GET("/trailer", h.getTrailer)
	}
}
