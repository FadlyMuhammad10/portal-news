package handler

import (
	"portalnews/internal/adapter/handler/request"
	"portalnews/internal/adapter/handler/response"
	"portalnews/internal/core/domain/entity"
	"portalnews/internal/core/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

var err error
var code string
var errResp response.ErrorResponseDefault
var validate = validator.New()

type AuthHandler interface {
	Login(c fiber.Ctx) error
}

type authHandler struct {
	authService service.AuthService
}

// Login implements [AuthHandler].
func (a *authHandler) Login(c fiber.Ctx) error {
	req := request.LoginRequest{}
	res := response.SuccessAuthResponse{}

	if err = c.Bind().Body(&req); err != nil {
		code = "[HANDLER Login -1]"
		log.Errorw(code, err)
		errResp.Meta.Status = false
		errResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errResp)
	}

	if err := validate.Struct(req); err != nil {
		code = "[HANDLER Login -2]"
		log.Errorw(code, err)
		errResp.Meta.Status = false
		errResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errResp)
	}

	reqLogin := entity.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	result, err := a.authService.GetUserByEmail(c.Context(), reqLogin)
	if err != nil {
		code = "[HANDLER Login -3]"
		log.Errorw(code, err)
		errResp.Meta.Status = false
		errResp.Meta.Message = err.Error()

		if err.Error() == "invalid password" {
			return c.Status(fiber.StatusUnauthorized).JSON(errResp)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(errResp)
	}

	res.Meta.Status = true
	res.Meta.Message = "Login successfull"
	res.AccessToken = result.TokenAccess
	res.ExpiresAt = result.ExpiresAt

	return c.JSON(res)
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}
