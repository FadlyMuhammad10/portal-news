package service

import (
	"context"
	"errors"
	"portalnews/config"
	"portalnews/internal/adapter/repository"
	"portalnews/internal/core/domain/entity"
	"portalnews/lib/auth"
	"portalnews/lib/conv"
	"time"

	"github.com/gofiber/fiber/v3/log"
	"github.com/golang-jwt/jwt/v5"
)

var err error
var code string

type AuthService interface {
	GetUserByEmail(ctx context.Context, req entity.LoginRequest) (*entity.AccessToken, error)
}

type authService struct {
	authRepo repository.AuthRepository
	cfg      *config.Config
	jwtToken auth.Jwt
}

// GetUserByEmail implements [AuthService].
func (a *authService) GetUserByEmail(ctx context.Context, req entity.LoginRequest) (*entity.AccessToken, error) {
	result, err := a.authRepo.GetUserByEmail(ctx, req)
	if err != nil {
		code = "[SERVICE] GetUserByEmail -1"
		log.Errorw(code, err)
		return nil, err
	}

	if checkPass := conv.CheckPasswordHash(req.Password, result.Password); !checkPass {
		code = "[SERVICE] GetUserByEmail -2"
		err = errors.New("invalid password")
		log.Errorw(code, err)
		return nil, err
	}

	jwtData := entity.JwtData{
		UserId: float64(result.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			ID:        string(result.ID),
		},
	}

	accToken, expiresAt, err := a.jwtToken.GenerateToken(&jwtData)
	if err != nil {
		code = "[SERVICE] GetUserByEmail -3"
		log.Errorw(code, err)
		return nil, err
	}

	resp := entity.AccessToken{
		TokenAccess: accToken,
		ExpiresAt:   expiresAt,
	}

	return &resp, nil
}

func NewAuthservice(authRepo repository.AuthRepository, cfg *config.Config, jwtToken auth.Jwt) AuthService {
	return &authService{authRepo: authRepo,
		cfg:      cfg,
		jwtToken: jwtToken}
}
