package authentication

type User struct {
	ID       int
	Password string `json:"password"`
}
