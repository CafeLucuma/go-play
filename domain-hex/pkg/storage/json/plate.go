package json

import "time"

type Plate struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
}
