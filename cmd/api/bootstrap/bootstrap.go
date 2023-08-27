package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	usuario "github.com/adnicolas/golang-hexagonal/internal"
	"github.com/adnicolas/golang-hexagonal/internal/creating"
	"github.com/adnicolas/golang-hexagonal/internal/fetching"
	"github.com/adnicolas/golang-hexagonal/internal/increasing"
	"github.com/adnicolas/golang-hexagonal/internal/platform/bus/inmemory"
	"github.com/adnicolas/golang-hexagonal/internal/platform/server"
	pg "github.com/adnicolas/golang-hexagonal/internal/platform/storage/postgresql"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

func Run() error {
	// envconfig library allows the use of environment variables for configuration
	var cfg config
	err := envconfig.Process("gohex", &cfg)
	if err != nil {
		return err
	}

	postgresURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		queryBus   = inmemory.NewQueryBus()
		eventBus   = inmemory.NewEventBus()
	)

	userRepository := pg.NewUserRepository(db, cfg.DbTimeout)

	creatingUserService := creating.NewUserService(userRepository, eventBus)
	fetchingUserService := fetching.NewUserService(userRepository)
	increasingUserService := increasing.NewUserCounterIncreaserService()

	createUserCommandHandler := creating.NewUserCommandHandler(creatingUserService)
	findUsersQueryHandler := fetching.NewUserQueryHandler(fetchingUserService)

	commandBus.RegisterCommand(creating.UserCommandType, createUserCommandHandler)
	queryBus.RegisterQuery(fetching.UserQueryType, findUsersQueryHandler)

	eventBus.Subscribe(usuario.UserCreatedEventType, creating.NewIncreaseUsersCounterOnUserCreated(increasingUserService))

	// Server initialization with empty context
	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus, queryBus)

	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `default:"localhost"`
	Port            uint          `default:"8081"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DbUser     string        `default:"postgres"`
	DbPassword string        `default:"dockerized_metadata"`
	DbHost     string        `default:"localhost"`
	DbPort     uint          `default:"5432"`
	DbName     string        `default:"metadata"`
	DbTimeout  time.Duration `default:"5s"`
}
