package pg

import (
	"context"
	"database/sql"
	"fmt"

	usuario "github.com/adnicolas/golang-hexagonal/internal"
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
		ID:       user.GetID().String(),
		Name:     user.GetName(),
		Surname:  user.GetSurname(),
		Password: user.GetPassword(),
		Email:    user.GetEmail(),
	}).Build()
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist user on database: %v", err)
	}

	return nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]usuario.GetUsersDto, error) {
	userSQLStruct := sqlbuilder.NewStruct(new(SqlUser)).WithTag("select").For(sqlbuilder.PostgreSQL)
	query, args := userSQLStruct.SelectFrom(sqlUserTable).Build()
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []usuario.GetUsersDto
	for rows.Next() {
		var item usuario.GetUsersDto
		if err := rows.Scan(&item.Id, &item.Name, &item.Surname, &item.Email); err != nil {
			return users, err
		}
		users = append(users, item)
	}
	if err = rows.Err(); err != nil {
		return users, fmt.Errorf("error trying to query users from database: %v", err)
	}
	return users, nil
}
