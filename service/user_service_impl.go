package service

import (
	"context"
	"fmt"
	"mfahmii/golang-restful/app"
	"mfahmii/golang-restful/exception"
	"mfahmii/golang-restful/helper"
	"mfahmii/golang-restful/model/domain"
	"mfahmii/golang-restful/model/web"
	"mfahmii/golang-restful/repository"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *app.Validation
	Config         *app.Config
}

func NewUserService(userRepository repository.UserRepository, DB *gorm.DB, validate *app.Validation, config *app.Config) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
		Config:         config,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Name: request.Name,
	}

	user, err = service.UserRepository.Save(ctx, tx, user)
	if err != nil {
		panic(err)
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Name = request.Name

	user = service.UserRepository.Update(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRepository.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {
	fmt.Println(userId)
	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindByEmail(ctx context.Context, email string) web.UserResponse {
	fmt.Println("test")
	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)

	return helper.ToUserResponses(users)
}
