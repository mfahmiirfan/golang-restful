package service

import (
	"context"
	"mfahmii/golang-restful/app"
	"mfahmii/golang-restful/exception"
	"mfahmii/golang-restful/helper"
	"mfahmii/golang-restful/model/domain"
	"mfahmii/golang-restful/model/web"
	"mfahmii/golang-restful/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *app.Validation
	Config         *app.Config
}

func NewAuthService(userRepository repository.UserRepository, DB *gorm.DB, validate *app.Validation, config *app.Config) AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
		Config:         config,
	}
}

func (service *AuthServiceImpl) SignUp(ctx context.Context, request web.UserSignUpRequest) web.UserResponse {
	service.Validate.RegisterStructValidation(func(sl validator.StructLevel) {
		request := sl.Current().Interface().(web.UserSignUpRequest)

		if request.PasswordConfirm != request.Password {
			sl.ReportError(request.PasswordConfirm, "passwordConfirm", "PasswordConfirm", "passwordConfirm", "")
		}
	}, web.UserSignUpRequest{})
	service.Validate.AddTranslation("passwordConfirm", "Passwords do not match")

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	user, err = service.UserRepository.Save(ctx, tx, user)
	if err != nil {
		panic(err)
	}

	return helper.ToUserResponse(user)
}

func (service *AuthServiceImpl) SignIn(ctx context.Context, request web.UserSignInRequest) web.TokenResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(exception.NewLoginError("Invalid email or Password"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		panic(exception.NewLoginError("Invalid email or Password"))
	}

	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)

	claims["sub"] = user.ID
	claims["role"] = "test"
	claims["exp"] = now.Add(service.Config.JwtExpiresIn).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte(service.Config.JwtSecret))
	helper.PanicIfError(err)

	return web.TokenResponse{
		Token: tokenString,
	}
}
