package repository

import (
	"context"
	"mfahmii/golang-restful/model/domain"

	"gorm.io/gorm"
)

type GiftRepository interface {
	Save(ctx context.Context, tx *gorm.DB, gift domain.Gift) domain.Gift
	Update(ctx context.Context, tx *gorm.DB, gift domain.Gift) domain.Gift
	Delete(ctx context.Context, tx *gorm.DB, gift domain.Gift)
	FindById(ctx context.Context, tx *gorm.DB, giftId int) (domain.Gift, error)
	FindAll(ctx context.Context, tx *gorm.DB) []domain.Gift
}
