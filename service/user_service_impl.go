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

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
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
	service.Validate.RegisterStructValidation(func(sl validator.StructLevel) {
		request := sl.Current().Interface().(web.UserCreateRequest)

		if request.PasswordConfirm != request.Password {
			sl.ReportError(request.PasswordConfirm, "passwordConfirm", "PasswordConfirm", "passwordConfirm", "")
		}
	}, web.UserCreateRequest{})
	service.Validate.AddTranslation("passwordConfirm", "Passwords do not match")
	fmt.Println(request)

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	fmt.Println(request)
	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
		Role:     &request.Role,
		Verified: &request.Verified,
	}

	fmt.Println(user)
	user, err = service.UserRepository.Save(ctx, tx, user)
	if err != nil {
		panic(err)
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	service.Validate.RegisterStructValidation(func(sl validator.StructLevel) {
		request := sl.Current().Interface().(web.UserUpdateRequest)

		if request.PasswordConfirm != request.Password {
			sl.ReportError(request.PasswordConfirm, "passwordConfirm", "PasswordConfirm", "passwordConfirm", "")
		}
	}, web.UserUpdateRequest{})
	service.Validate.AddTranslation("passwordConfirm", "Passwords do not match")

	fmt.Println(request)

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashedPassword)
	user.Role = &request.Role
	user.Verified = &request.Verified

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
