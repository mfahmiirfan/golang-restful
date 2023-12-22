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
	// SQL := "insert into category(name) values (?)"
	// result, err := tx.ExecContext(ctx, SQL, category.Name)
	response := tx.WithContext(ctx).Create(&category)
	helper.PanicIfError(response.Error)

	// id, err := result.LastInsertId()
	// helper.PanicIfError(err)

	// category.Id = int(id)
	fmt.Println(category.Id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, category domain.Category) domain.Category {
	// SQL := "update category set name = ? where id = ?"
	// _, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	err := tx.WithContext(ctx).Save(&category).Error
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, category domain.Category) {
	// SQL := "delete from category where id = ?"
	// _, err := tx.ExecContext(ctx, SQL, category.Id)
	err := tx.WithContext(ctx).Delete(&category).Error
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, categoryId int) (domain.Category, error) {
	// SQL := "select id, name from category where id = ?"
	// rows, err := tx.QueryContext(ctx, SQL, categoryId)
	category := domain.Category{}
	err := tx.WithContext(ctx).First(&category, categoryId).Error
	fmt.Println("ini error:")
	fmt.Println(reflect.TypeOf(err))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("masuk error record not found")
			return category, errors.New("category is not found")
		}
		helper.PanicIfError(err)
	}
	// defer rows.Close()

	// category := domain.Category{}
	// if rows.Next() {
	// 	err := rows.Scan(&category.Id, &category.Name)
	// 	helper.PanicIfError(err)
	// 	return category, nil
	// } else {
	// 	return category, errors.New("category is not found")
	// }
	return category, nil
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []domain.Category {
	// SQL := "select id, name from category"
	// rows, err := tx.QueryContext(ctx, SQL)
	categories := []domain.Category{}
	err := tx.WithContext(ctx).Find(&categories).Error
	helper.PanicIfError(err)
	// defer rows.Close()

	// for rows.Next() {
	// 	category := domain.Category{}
	// 	err := rows.Scan(&category.Id, &category.Name)
	// 	helper.PanicIfError(err)
	// 	categories = append(categories, category)
	// }
	return categories
}
