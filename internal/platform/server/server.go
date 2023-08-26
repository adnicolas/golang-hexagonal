package server

import (
	"fmt"
	"log"

	"github.com/adnicolas/golang-hexagonal/internal/platform/server/handler/health"
	"github.com/adnicolas/golang-hexagonal/internal/platform/server/handler/users"
	"github.com/adnicolas/golang-hexagonal/kit/bus"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine   *gin.Engine
	httpAddr string
	// deps
	bus bus.Bus
}

// Gin wrapper
func New(host string, port uint, myBus bus.Bus) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		bus:      myBus,
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/user", users.CreateController(s.bus))
	s.engine.GET("/users", users.FindAllController(s.bus))
}
