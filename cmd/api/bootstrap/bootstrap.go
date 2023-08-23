package bootstrap

import "github.com/adnicolas/golang-hexagonal/internal/platform/server"

const (
	host = "localhost"
	port = 8081
)

func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}
