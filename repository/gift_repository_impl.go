package repository

import (
	"context"
	"errors"
	"fmt"
	"mfahmii/golang-restful/helper"
	"mfahmii/golang-restful/model/domain"
	"reflect"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, category domain.Category) domain.Category {
	err := tx.WithContext(ctx).Create(&category).Error
	helper.PanicIfError(err)

	fmt.Println(category.Id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, category domain.Category) domain.Category {
	err := tx.WithContext(ctx).Save(&category).Error
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, category domain.Category) {
	err := tx.WithContext(ctx).Delete(&category).Error
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, categoryId int) (domain.Category, error) {
	category := domain.Category{}
	err := tx.WithContext(ctx).First(&category, categoryId).Error
	fmt.Println("ini error:")
	fmt.Println(reflect.TypeOf(err))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("masuk error record not found")
			return category, errors.New("category is not found")
		}
		panic(err)
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []domain.Category {
	categories := []domain.Category{}
	err := tx.WithContext(ctx).Find(&categories).Error
	helper.PanicIfError(err)

	return categories
}
