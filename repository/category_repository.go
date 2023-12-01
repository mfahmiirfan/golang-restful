package repository

import (
	"context"
	"mfahmii/golang-restful/model/domain"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *gorm.DB, category domain.Category) domain.Category
	Update(ctx context.Context, tx *gorm.DB, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *gorm.DB, category domain.Category)
	FindById(ctx context.Context, tx *gorm.DB, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *gorm.DB) []domain.Category
}
