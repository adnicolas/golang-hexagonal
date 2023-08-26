package pg

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	usuario "github.com/adnicolas/golang-hexagonal/internal"
	"github.com/huandu/go-sqlbuilder"
)

// UserRepository is a PostgreSQL usuario.UserRepository implementation
type UserRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewUserRepository initializes a PostgreSQL-based implementation of usuario.UserRepository
func NewUserRepository(db *sql.DB, dbTimeout time.Duration) *UserRepository {
	return &UserRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Save implements the usuario.UserRepository interface
func (r *UserRepository) Save(ctx context.Context, user usuario.User) error {
	userSQLStruct := sqlbuilder.NewStruct(new(sqlUser)).For(sqlbuilder.PostgreSQL)
	query, args := userSQLStruct.InsertInto(sqlUserTable, sqlUser{
		ID:       user.GetID().String(),
		Name:     user.GetName(),
		Surname:  user.GetSurname(),
		Password: user.GetPassword(),
		Email:    user.GetEmail(),
	}).Build()

	// Context pattern example
	// A timeout is added for when trying to persist a user in the DB
	// It's necessary to pass the parent context to the new context
	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist user on database: %v", err)
	}

	return nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]usuario.GetUsersDto, error) {
	userSQLStruct := sqlbuilder.NewStruct(new(sqlUser)).WithTag("select").For(sqlbuilder.PostgreSQL)
	query, args := userSQLStruct.SelectFrom(sqlUserTable).Build()

	// Context pattern example
	// A timeout is added for when trying to fetch users from DB
	// It's necessary to pass the parent context to the new context
	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return []usuario.GetUsersDto{}, fmt.Errorf("error trying to query users from database: %v", err)
	}
	defer rows.Close()
	var users []usuario.GetUsersDto
	for rows.Next() {
		var item usuario.GetUsersDto
		if err := rows.Scan(&item.Id, &item.Name, &item.Surname, &item.Email); err != nil {
			return []usuario.GetUsersDto{}, fmt.Errorf("error trying to query users from database: %v", err)
		}
		users = append(users, item)
	}
	if err = rows.Err(); err != nil {
		return users, fmt.Errorf("error trying to query users from database: %v", err)
	}
	return users, nil
}
