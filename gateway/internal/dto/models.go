package dto

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserEntity struct {
	Login    string
	Password string
}

type LoginResponse struct {
	Token string `json:"token"`
}
type Credentials struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}
