package entity

import "github.com/golang-jwt/jwt/v5"

// payloadnya informasi jwt
type JwtData struct {
	UserId               float64 `json:"user_id"`
	jwt.RegisteredClaims         // ini sama untuk menyimpan claims yang nanti menjadi payload jwt
}
