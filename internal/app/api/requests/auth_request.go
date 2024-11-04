package requests

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthRefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
