package users

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

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

type LoginData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponseUser struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type LoginResponse struct {
	Success bool              `json:"success"`
	Error   string            `json:"error"`
	User    LoginResponseUser `json:"user"`
}

type LogoutResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
