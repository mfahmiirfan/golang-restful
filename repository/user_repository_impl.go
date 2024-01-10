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

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error) {
	err := tx.WithContext(ctx).Create(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return user, errors.New("Duplicate entry")
		}
		panic(err)
	}

	fmt.Println(user.ID)
	return user, nil
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, user domain.User) domain.User {
	err := tx.WithContext(ctx).Save(&user).Error
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, user domain.User) {
	err := tx.WithContext(ctx).Delete(&user).Error
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, userId int) (domain.User, error) {
	user := domain.User{}
	err := tx.WithContext(ctx).First(&user, userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user is not found")
		}
		panic(err)
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *gorm.DB, email string) (domain.User, error) {
	user := domain.User{}
	err := tx.WithContext(ctx).First(&user, "email = ?", email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user is not found")
		}
		panic(err)
	}
	return user, nil
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []domain.User {
	users := []domain.User{}
	err := tx.WithContext(ctx).Find(&users).Error
	helper.PanicIfError(err)

	return users
}
