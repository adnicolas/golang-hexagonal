package server

import (
	"fmt"
	"log"

	"github.com/adnicolas/golang-hexagonal/internal/creating"
	"github.com/adnicolas/golang-hexagonal/internal/fetching"
	"github.com/adnicolas/golang-hexagonal/internal/platform/server/handler/health"
	"github.com/adnicolas/golang-hexagonal/internal/platform/server/handler/users"
	"github.com/adnicolas/golang-hexagonal/kit/command"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine   *gin.Engine
	httpAddr string
	// deps
	creatingUserService creating.UserService
	fetchingUserService fetching.UserService
	commandBus          command.Bus
}

// Gin wrapper
func New(host string, port uint, fetchingUserService fetching.UserService, creatingUserService creating.UserService, commandBus command.Bus) Server {
	srv := Server{
		engine:              gin.New(),
		httpAddr:            fmt.Sprintf("%s:%d", host, port),
		fetchingUserService: fetchingUserService,
		creatingUserService: creatingUserService,
		commandBus:          commandBus,
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
	s.engine.POST("/user", users.CreateController(s.commandBus))
	s.engine.GET("/users", users.FindAllController(s.fetchingUserService))
}
