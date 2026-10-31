package middleware

import (
	"portalnews/config"
	"portalnews/internal/adapter/handler/response"
	"portalnews/lib/auth"
	"strings"

	"github.com/gofiber/fiber/v3"
)

type Middleware interface {
	CheckToken() fiber.Handler
}

type Options struct {
	authJwt auth.Jwt
}

// CheckToken implements [Middleware].
func (o *Options) CheckToken() fiber.Handler {
	return func(c fiber.Ctx) error {
		var errorResponse response.ErrorResponseDefault
		// cek header
		authHandler := c.Get("Authorization")
		if authHandler == "" {
			errorResponse.Status = false
			errorResponse.Message = "Missing Authorization header"
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)
		}

		safeTokenString := strings.TrimSpace(strings.Split(authHandler, "Bearer")[1])
		claims, err := o.authJwt.VerifyToken(safeTokenString)
		if err != nil {
			errorResponse.Status = false
			errorResponse.Message = "Invalid Token"
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)
		}
		c.Locals("user", claims) // disimpan di local mirip seperti localstorage/cookies
		return c.Next()
	}
}

func NewMiddleware(cfg *config.Config) Middleware {
	opt := new(Options)
	opt.authJwt = auth.NewJwt(cfg)

	return opt
}
