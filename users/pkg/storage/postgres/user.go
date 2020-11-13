package postgres

import "time"

type User struct {
	ID        int
	Name      string
	LastName  string
	Password  string
	Email     string
	IsAdmin   bool
	CreatedOn time.Time
}
