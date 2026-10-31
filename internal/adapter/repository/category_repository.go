package repository

import (
	"context"
	"errors"
	"portalnews/internal/core/domain/entity"
	"portalnews/internal/core/domain/model"

	"github.com/gofiber/fiber/v3/log"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategories(ctx context.Context) ([]entity.CategoryEntity, error)
	GetCategoryById(ctx context.Context, id int64) (*entity.CategoryEntity, error) // pointer agar null jika ada error
	CreateCategory(ctx context.Context, req entity.CategoryEntity) (int64, error)
	EditCategoryById(ctx context.Context, req entity.CategoryEntity) error
	DeleteCategoryById(ctx context.Context, id int64) error
}

type categoryRepository struct {
	db *gorm.DB
}

// CreateCategory implements [CategoryRepository].
func (c *categoryRepository) CreateCategory(ctx context.Context, req entity.CategoryEntity) (int64, error) {
	category := model.Category{
		Title:       req.Title,
		Slug:        req.Slug,
		CreatedByID: req.User.ID,
	}

	if err := c.db.WithContext(ctx).Create(&category).Error; err != nil {
		code = "[REPOSITORY] CreateCategory -1"
		log.Errorw(code, err)
		return 0, err
	}

	return category.ID, nil
}

// DeleteCategoryById implements [CategoryRepository].
func (c *categoryRepository) DeleteCategoryById(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// EditCategoryById implements [CategoryRepository].
func (c *categoryRepository) EditCategoryById(ctx context.Context, req entity.CategoryEntity) error {
	panic("unimplemented")
}

// GetCategories implements [CategoryRepository].
func (c *categoryRepository) GetCategories(ctx context.Context) ([]entity.CategoryEntity, error) {
	var modelCategories []model.Category

	err := c.db.Order("created_at DESC").Preload("User").Find(&modelCategories).Error
	if err != nil {
		code = "[REPOSITORY] GetCategories -1"
		log.Errorw(code, err)
		return nil, err
	}

	if len(modelCategories) == 0 {
		code = "[REPOSITORY] GetCategories -2"
		log.Errorw(code, err)
		err = errors.New("data not found")
		return nil, err
	}

	var resps []entity.CategoryEntity
	for _, val := range modelCategories {
		resps = append(resps, entity.CategoryEntity{
			ID:    val.ID,
			Title: val.Title,
			Slug:  val.Slug,
			User: entity.UserEntity{
				ID:    val.User.ID,
				Name:  val.User.Name,
				Email: val.User.Email,
			},
		})
	}
	return resps, nil
}

// GetCategoryById implements [CategoryRepository].
func (c *categoryRepository) GetCategoryById(ctx context.Context, id int64) (*entity.CategoryEntity, error) {
	var modelCategories model.Category

	err = c.db.Where("id = ?", id).Preload("User").First(&modelCategories).Error

	if err != nil {
		code = "[REPOSITORY] GetUserByEmail -1"
		log.Errorw(code, err)
		return nil, err
	}

	resp := entity.CategoryEntity{
		ID:    modelCategories.ID,
		Title: modelCategories.Title,
		Slug:  modelCategories.Slug,
		User: entity.UserEntity{
			ID:    modelCategories.User.ID,
			Name:  modelCategories.User.Name,
			Email: modelCategories.User.Email,
		},
	}

	return &resp, nil
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}
