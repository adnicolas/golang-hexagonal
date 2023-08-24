package pg

import (
	"context"
	"database/sql"
	"fmt"

	usuario "github.com/adnicolas/golang-hexagonal/internal/platform"
	"github.com/huandu/go-sqlbuilder"
)

// UserRepository is a PostgreSQL usuario.UserRepository implementation
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository initializes a PostgreSQL-based implementation of usuario.UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Save implements the usuario.UserRepository interface
func (r *UserRepository) Save(ctx context.Context, user usuario.User) error {
	userSQLStruct := sqlbuilder.NewStruct(new(SqlUser)).For(sqlbuilder.PostgreSQL)
	query, args := userSQLStruct.InsertInto(sqlUserTable, SqlUser{
		ID:       user.ID(),
		Name:     user.Name(),
		Surname:  user.Surname(),
		Password: user.Password(),
		Email:    user.Email(),
	}).Build()
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist user on database: %v", err)
	}

	return nil
}
