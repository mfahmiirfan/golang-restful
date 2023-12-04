package repository

import (
	"context"
	"mfahmii/golang-restful/model/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, tx *gorm.DB, user domain.User) domain.User
	Update(ctx context.Context, tx *gorm.DB, user domain.User) domain.User
	Delete(ctx context.Context, tx *gorm.DB, user domain.User)
	FindById(ctx context.Context, tx *gorm.DB, userId uuid.UUID) (domain.User, error)
	FindAll(ctx context.Context, tx *gorm.DB) []domain.User
}
