package server

import (
	"fmt"
	"log"

	usuario "github.com/adnicolas/golang-hexagonal/internal"
	"github.com/adnicolas/golang-hexagonal/internal/platform/server/handler/health"
	"github.com/adnicolas/golang-hexagonal/internal/platform/server/handler/users"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine   *gin.Engine
	httpAddr string
	// deps
	userRepository usuario.UserRepository
}

// Gin wrapper
func New(host string, port uint, userRepository usuario.UserRepository) Server {
	srv := Server{
		engine:         gin.New(),
		httpAddr:       fmt.Sprintf("%s:%d", host, port),
		userRepository: userRepository,
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
	s.engine.POST("/user", users.CreateSaveController(s.userRepository))
	s.engine.GET("/users", users.CreateFindAllController(s.userRepository))
}
