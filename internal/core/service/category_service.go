package service

import (
	"context"
	"portalnews/internal/adapter/repository"
	"portalnews/internal/core/domain/entity"

	"github.com/gofiber/fiber/v3/log"
)

type CategoryService interface {
	GetCategories(ctx context.Context) ([]entity.CategoryEntity, error)
	GetCategoryById(ctx context.Context, id int64) (*entity.CategoryEntity, error) // pointer agar null jika ada error
	CreateCategory(ctx context.Context, req entity.CategoryEntity) (int64, error)
	EditCategoryById(ctx context.Context, req entity.CategoryEntity) error
	DeleteCategoryById(ctx context.Context, id int64) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

// CreateCategory implements [CategoryService].
func (c *categoryService) CreateCategory(ctx context.Context, req entity.CategoryEntity) (int64, error) {
	createCategory, err := c.categoryRepo.CreateCategory(ctx, req)
	if err != nil {
		code = "[SERVICE] CreateCategory -1"
		log.Errorw(code, err)
		return 0, err
	}

	return createCategory, nil
}

// DeleteCategoryById implements [CategoryService].
func (c *categoryService) DeleteCategoryById(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// EditCategoryById implements [CategoryService].
func (c *categoryService) EditCategoryById(ctx context.Context, req entity.CategoryEntity) error {
	panic("unimplemented")
}

// GetCategoryById implements [CategoryService].
func (c *categoryService) GetCategoryById(ctx context.Context, id int64) (*entity.CategoryEntity, error) {
	result, err := c.categoryRepo.GetCategoryById(ctx, id)
	if err != nil {
		code = "[SERVICE] GetCategoryById -1"
		log.Errorw(code, err)
		return nil, err
	}

	return result, nil
}

// GetCategories implements [CategoryService].
func (c *categoryService) GetCategories(ctx context.Context) ([]entity.CategoryEntity, error) {
	result, err := c.categoryRepo.GetCategories(ctx)
	if err != nil {
		code = "[SERVICE] GetCategories -1"
		log.Errorw(code, err)
		return nil, err
	}

	return result, nil
}

func NewCategoryervice(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}
