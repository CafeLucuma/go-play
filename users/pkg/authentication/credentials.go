package authentication

type LoginCredentials struct {
	Password *string `json:"password"`
	Email    *string `json:"email"`
}
