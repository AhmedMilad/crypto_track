package utils

type SignUpRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type Response struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    any    `json:"data"`
}
