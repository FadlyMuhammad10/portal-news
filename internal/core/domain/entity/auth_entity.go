package entity

type LoginRequest struct {
	Email    string
	Password string
}

type AccessToken struct {
	TokenAccess string
	ExpiresAt   int64
}
