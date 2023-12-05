package repository

import (
	"context"
	"errors"
	"fmt"
	"mfahmii/golang-restful/helper"
	"mfahmii/golang-restful/model/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, user domain.User) domain.User {
	// SQL := "insert into category(name) values (?)"
	// result, err := tx.ExecContext(ctx, SQL, category.Name)
	response := tx.WithContext(ctx).Create(&user)
	helper.PanicIfError(response.Error)

	// id, err := result.LastInsertId()
	// helper.PanicIfError(err)

	// category.Id = int(id)
	fmt.Println(user.ID)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, user domain.User) domain.User {
	// SQL := "update category set name = ? where id = ?"
	// _, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	err := tx.WithContext(ctx).Save(&user).Error
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, user domain.User) {
	// SQL := "delete from category where id = ?"
	// _, err := tx.ExecContext(ctx, SQL, category.Id)
	err := tx.WithContext(ctx).Delete(&user).Error
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, userId int) (domain.User, error) {
	// SQL := "select id, name from category where id = ?"
	// rows, err := tx.QueryContext(ctx, SQL, categoryId)
	var user domain.User
	err := tx.WithContext(ctx).First(&user, userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user is not found")
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
	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *gorm.DB, email string) (domain.User, error) {
	var user domain.User
	err := tx.WithContext(ctx).First(&user, "email = ?", email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user is not found")
		}
		helper.PanicIfError(err)
	}
	return user, nil
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []domain.User {
	// SQL := "select id, name from category"
	// rows, err := tx.QueryContext(ctx, SQL)
	var categories []domain.User
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

func (repository *UserRepositoryImpl) test() error {
	return nil
}
