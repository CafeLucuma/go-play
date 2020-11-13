package adding

type User struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
}
