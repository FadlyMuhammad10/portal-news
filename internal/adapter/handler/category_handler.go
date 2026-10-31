package handler

import (
	"portalnews/internal/adapter/handler/request"
	"portalnews/internal/adapter/handler/response"
	"portalnews/internal/core/domain/entity"
	"portalnews/internal/core/service"
	"portalnews/lib/conv"
	validatorLib "portalnews/lib/validator"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

var defaultSuccessResponse response.SuccessResponseDefault
var defaultCreateSuccessResponse response.SuccessCreateResponseDefault

type CategoryHandler interface {
	GetCategories(c fiber.Ctx) error
	GetCategoryById(c fiber.Ctx) error
	CreateCategory(c fiber.Ctx) error
	EditCategoryById(c fiber.Ctx) error
	DeleteCategoryById(c fiber.Ctx) error
}

type categoryHandler struct {
	categoryService service.CategoryService
}

// CreateCategory implements [CategoryHandler].
func (ch *categoryHandler) CreateCategory(c fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	userID := claims.UserId
	if userID == 0 {
		code = "[HANDLER CreateCategory -1]"
		log.Errorw(code, err)
		errResp.Meta.Status = false
		errResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusUnauthorized).JSON(errResp)
	}
	req := request.CategoryCreateRequest{}
	if err = c.Bind().Body(&req); err != nil {
		code = "[HANDLER CreateCategory -2]"
		log.Errorw(code, err)
		errResp.Meta.Status = false
		errResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errResp)
	}

	if err := validatorLib.ValidateStruct(req); err != nil {
		code = "[HANDLER CreateCategory -3]"
		log.Errorw(code, err)
		errResp.Meta.Status = false
		errResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errResp)
	}

	slugGenerate := conv.GenerateSlug(req.Title)

	reqCreateCategory := entity.CategoryEntity{
		Title: req.Title,
		Slug:  slugGenerate,
		User: entity.UserEntity{
			ID: int64(userID),
		},
	}

	categoryID, err := ch.categoryService.CreateCategory(c.Context(), reqCreateCategory)
	if err != nil {
		code = "[HANDLER CreateCategory -4]"
		log.Errorw(code, err)

		errResp.Meta.Status = false
		errResp.Meta.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(errResp)
	}

	createCategoryResp := response.SuccessCreateCategory{
		ID: categoryID,
	}

	defaultCreateSuccessResponse.Meta.Status = true
	defaultCreateSuccessResponse.Meta.Message = "Success create category"
	defaultCreateSuccessResponse.Data = createCategoryResp

	return c.Status(fiber.StatusCreated).JSON(defaultCreateSuccessResponse)
}

// DeleteCategoryById implements [CategoryHandler].
func (ch *categoryHandler) DeleteCategoryById(c fiber.Ctx) error {
	panic("unimplemented")
}

// EditCategoryById implements [CategoryHandler].
func (ch *categoryHandler) EditCategoryById(c fiber.Ctx) error {
	panic("unimplemented")
}

// GetCategories implements [CategoryHandler].
func (ch *categoryHandler) GetCategories(c fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	userID := claims.UserId
	if userID == 0 {
		code = "[HANDLER GetCategories -1]"
		log.Errorw(code, err)
		errResp.Meta.Status = false
		errResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusUnauthorized).JSON(errResp)
	}

	results, err := ch.categoryService.GetCategories(c.Context())
	if err != nil {
		code = "[HANDLER GetCategories -2]"
		log.Errorw(code, err)
		errResp.Meta.Status = false
		errResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusNotFound).JSON(errResp)
	}

	categoriesResps := []response.SuccessCategoryResponse{}
	for _, result := range results {
		categoriesResp := response.SuccessCategoryResponse{
			ID:          result.ID,
			Title:       result.Title,
			Slug:        result.Slug,
			CreatedName: result.User.Name,
		}
		categoriesResps = append(categoriesResps, categoriesResp)
	}
	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Success get categories"
	defaultSuccessResponse.Data = categoriesResps

	return c.Status(fiber.StatusOK).JSON(defaultSuccessResponse)
}

// GetCategoryById implements [CategoryHandler].
func (ch *categoryHandler) GetCategoryById(c fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	userID := claims.UserId
	if userID == 0 {
		code = "[HANDLER GetCategoryById -1]"
		log.Errorw(code, err)
		errResp.Meta.Status = false
		errResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusUnauthorized).JSON(errResp)
	}

	id := c.Params("id")
	parseParamsId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		code = "[HANDLER GetCategoryById -2]"
		log.Errorw(code, err)
		errResp.Meta.Status = false
		errResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errResp)
	}

	result, err := ch.categoryService.GetCategoryById(c.Context(), parseParamsId)
	if err != nil {
		code = "[HANDLER GetCategoryById -3]"
		log.Errorw(code, err)
		errResp.Meta.Status = false
		errResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusNotFound).JSON(errResp)
	}

	categoriesResp := response.SuccessCategoryResponse{
		ID:          result.ID,
		Title:       result.Title,
		Slug:        result.Slug,
		CreatedName: result.User.Name,
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Success get category by id"
	defaultSuccessResponse.Data = categoriesResp

	return c.Status(fiber.StatusOK).JSON(defaultSuccessResponse)
}

func NewCateoryHandler(categoryService service.CategoryService) CategoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
	}
}
