package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/adnicolas/golang-hexagonal/internal/creating"
	"github.com/adnicolas/golang-hexagonal/internal/platform/server"
	pg "github.com/adnicolas/golang-hexagonal/internal/platform/storage/postgresql"
	_ "github.com/lib/pq"
)

const (
	host       = "localhost"
	port       = 8081
	dbUser     = "postgres"
	dbPassword = "dockerized_metadata"
	dbHost     = "localhost"
	dbPort     = 5433
	dbName     = "metadata"
)

func Run() error {
	postgresURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		return err
	}
	userRepository := pg.NewUserRepository(db)
	creatingCourseService := creating.NewUserService(userRepository)
	// Server initialization
	srv := server.New(host, port, userRepository, creatingCourseService)
	return srv.Run()
}
