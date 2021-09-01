package users

type RegisterData struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterResponseUser struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type RegisterResponse struct {
	Success bool                 `json:"success"`
	Error   string               `json:"error"`
	User    RegisterResponseUser `json:"user"`
}
