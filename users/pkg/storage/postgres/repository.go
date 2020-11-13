package postgres

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	"github.com/CafeLucuma/go-play/users/pkg/adding"
	"github.com/CafeLucuma/go-play/users/pkg/authentication"
	"github.com/CafeLucuma/go-play/users/pkg/logging"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

// NewStorage returns a new JSON  storage
func NewStorage() (*Storage, error) {
	storage := new(Storage)

	dbURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		logging.Debug.Printf("env variable %s empty", "DATABASE_URL")
		logging.Error.Printf("%s", "Cant load databse url from environment")
		return nil, errors.New("Cant load databse url from environment")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	storage.db = db

	if err := storage.db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return storage, nil
}

const (
	SELECT_USER_BY_EMAIL = "SELECT user_id, salted_password FROM users WHERE email = $1"
	INSERT_USER          = "INSERT INTO users (name, last_name, email, salted_password, is_admin, created_on) VALUES ($1, $2, $3, $4, $5, $6)"
)

func (s *Storage) CloseDB() {
	logging.Error.Printf("Closing db...")
	log.Fatal(s.db.Close())
}

func (s *Storage) GetUser(email string) (*authentication.User, error) {

	var user authentication.User
	userRow := s.db.QueryRow(SELECT_USER_BY_EMAIL, email)
	if err := userRow.Scan(&user.ID, &user.Password); err != nil {
		logging.Error.Printf("Error obtaining user with user email %v", email)
		return nil, err
	}

	return &user, nil
}

func (s *Storage) AddUser(u adding.User) error {

	logging.Info.Printf("Inserting user %+v to db", u)

	newUser := User{
		Name:      u.Name,
		LastName:  u.LastName,
		Password:  u.Password,
		Email:     u.Email,
		IsAdmin:   u.IsAdmin,
		CreatedOn: time.Now(),
	}

	_, insertErr := s.db.Exec(INSERT_USER, newUser.Name, newUser.LastName, newUser.Email, newUser.Password, newUser.IsAdmin, newUser.CreatedOn)
	if insertErr != nil {
		logging.Error.Printf("Cant insert new user to database: ", insertErr.Error())
		return insertErr
	}

	return nil
}
