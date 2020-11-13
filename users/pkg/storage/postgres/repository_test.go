package postgres

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestGetUserFound(t *testing.T) {
	db, mock := NewMock()
	repo := Storage{db: db}

	//defer repo.CloseDB()

	query := regexp.QuoteMeta(SELECT_USER_BY_EMAIL)
	rows := sqlmock.NewRows([]string{"user_id", "salted_password"}).AddRow(1, "asdasdasd")

	mock.ExpectQuery(query).WithArgs("ocm@mailinator.com").WillReturnRows(rows)

	user, err := repo.GetUser("ocm@mailinator.com")
	assert.NotNil(t, user)
	assert.Nil(t, err)
}

func TestGetUserNotFound(t *testing.T) {
	db, mock := NewMock()
	repo := Storage{db: db}

	//defer repo.CloseDB()

	query := regexp.QuoteMeta(SELECT_USER_BY_EMAIL)

	rows := sqlmock.NewRows([]string{"user_id", "salted_password"})

	mock.ExpectQuery(query).WithArgs("ocm@mailinator.com").WillReturnRows(rows)

	user, err := repo.GetUser("ocm@mailinator.com")
	assert.NotNil(t, err)
	assert.Nil(t, user)
}
