package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/adnicolas/golang-hexagonal/internal/creating"
	"github.com/adnicolas/golang-hexagonal/internal/fetching"
	"github.com/adnicolas/golang-hexagonal/internal/platform/bus/inmemory"
	"github.com/adnicolas/golang-hexagonal/internal/platform/server"
	pg "github.com/adnicolas/golang-hexagonal/internal/platform/storage/postgresql"
	_ "github.com/lib/pq"
)

const (
	host            = "localhost"
	port            = 8081
	shutdownTimeout = 10 * time.Second

	dbUser     = "postgres"
	dbPassword = "dockerized_metadata"
	dbHost     = "localhost"
	dbPort     = 5433
	dbName     = "metadata"
	dbTimeout  = 5 * time.Second
)

func Run() error {
	postgresURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		return err
	}

	userRepository := pg.NewUserRepository(db, dbTimeout)
	creatingUserService := creating.NewUserService(userRepository)
	fetchingUserService := fetching.NewUserService(userRepository)

	var (
		commandBus = inmemory.NewCommandBus()
		queryBus   = inmemory.NewQueryBus()
	)
	createUserCommandHandler := creating.NewUserCommandHandler(creatingUserService)
	commandBus.RegisterCommand(creating.UserCommandType, createUserCommandHandler)
	findUsersQueryHandler := fetching.NewUserQueryHandler(fetchingUserService)
	queryBus.RegisterQuery(fetching.UserQueryType, findUsersQueryHandler)

	// Server initialization with empty context
	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus, queryBus)

	return srv.Run(ctx)
}
