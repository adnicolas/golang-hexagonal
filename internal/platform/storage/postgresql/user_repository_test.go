package pg

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	usuario "github.com/adnicolas/golang-hexagonal/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_UserRepository_Save_RepositoryError(t *testing.T) {
	userId, userName, userSurname, userPassword, userEmail := "c226f125-7c63-4db6-ac47-aaf28baebeb5", "Adri", "Nico", "myPass", "adri@gmail.com"
	user, err := usuario.NewUser(userId, userName, userSurname, userPassword, userEmail)
	require.NoError(t, err)

	// sqlmock library allows to mock the DB connection
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO public.user (uuid, name, surname, password, email) VALUES ($1, $2, $3, $4, $5)").
		WithArgs(userId, userName, userSurname, userPassword, userEmail).
		WillReturnError(errors.New("something failed with user repository"))

	repo := NewUserRepository(db)

	err = repo.Save(context.Background(), user)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

// Happy path
func Test_UserRepository_Save_Succeed(t *testing.T) {
	userId, userName, userSurname, userPassword, userEmail := "c226f125-7c63-4db6-ac47-aaf28baebeb5", "Adri", "Nico", "myPass", "adri@gmail.com"
	user, err := usuario.NewUser(userId, userName, userSurname, userPassword, userEmail)
	require.NoError(t, err)

	// sqlmock library allows to mock the DB connection
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO public.user (uuid, name, surname, password, email) VALUES ($1, $2, $3, $4, $5)").
		WithArgs(userId, userName, userSurname, userPassword, userEmail).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewUserRepository(db)

	err = repo.Save(context.Background(), user)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}
