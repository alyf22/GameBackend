package Models

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
