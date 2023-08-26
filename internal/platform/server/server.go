package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/adnicolas/golang-hexagonal/internal/platform/server/handler/health"
	"github.com/adnicolas/golang-hexagonal/internal/platform/server/handler/users"
	"github.com/adnicolas/golang-hexagonal/kit/bus"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type Server struct {
	engine          *gin.Engine
	httpAddr        string
	shutdownTimeout time.Duration
	// deps
	bus bus.Bus
}

// Gin wrapper
func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, myBus bus.Bus) (context.Context, Server) {
	srv := Server{
		engine:          gin.New(),
		httpAddr:        fmt.Sprintf("%s:%d", host, port),
		shutdownTimeout: shutdownTimeout,
		bus:             myBus,
	}

	// Global logger middleware
	srv.engine.Use(gin.Logger())

	// Global recovery middleware recovers the service from any panics and logs a 500 error if one exists
	srv.engine.Use(gin.Recovery())

	srv.registerRoutes()
	return serverContext(ctx), srv
}

// The context is necessary to perform graceful shutdown
func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)

	// We run the native Go server (passing it Gin's server as a router in the Handler) so we can do graceful shutdown
	// see https://github.com/gin-gonic/gin/issues/2304

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shutdown", err)
		}
	}()

	<-ctx.Done()
	ctxShutdown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	// Go native graceful shutdown
	return srv.Shutdown(ctxShutdown)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/user", users.CreateController(s.bus))
	s.engine.GET("/users", users.FindAllController(s.bus))
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
