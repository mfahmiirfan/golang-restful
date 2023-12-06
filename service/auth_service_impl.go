package service

import (
	"fmt"
	"mfahmii/golang-restful/config"
	"mfahmii/golang-restful/exception"
	"mfahmii/golang-restful/helper"
	"mfahmii/golang-restful/model/web"
	"mfahmii/golang-restful/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
	Config         *config.Config
}

func NewAuthService(userRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate, config *config.Config) AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
		Config:         config,
	}
}

func (service *AuthServiceImpl) SignIn(ctx *fiber.Ctx, request web.UserSignInRequest) web.TokenResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx.Context(), tx, request.Email)
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
	claims["exp"] = now.Add(service.Config.JwtExpiresIn).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte(service.Config.JwtSecret))
	helper.PanicIfError(err)

	ctx.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   service.Config.JwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	return web.TokenResponse{
		Token: tokenString,
	}
}

func (service *AuthServiceImpl) SignUp(ctx *fiber.Ctx, request web.UserSignUpRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	if request.Password != request.PasswordConfirm {
		validationErrs := validator.ValidationErrors{}

		// Simulate adding custom errors manually
		validationErrs = append(validationErrs, &validator.FieldError{
			StructField: "Username",
			FieldError:  fmt.Errorf("Username is required"),
			Tag:         "required",
		})
	}
}
