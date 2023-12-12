package service

import (
	"mfahmii/golang-restful/app"
	"mfahmii/golang-restful/exception"
	"mfahmii/golang-restful/helper"
	"mfahmii/golang-restful/model/web"
	"mfahmii/golang-restful/repository"
	"time"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *app.Validation
	Config         *app.Config
	Trans          *ut.Translator
}

func NewAuthService(userRepository repository.UserRepository, DB *gorm.DB, validate *app.Validation, config *app.Config) AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
		Config:         config,
	}
}

func (service *AuthServiceImpl) SignUp(ctx *fiber.Ctx, request web.UserSignUpRequest) web.UserResponse {
	service.Validate.RegisterStructValidation(func(sl validator.StructLevel) {
		request := sl.Current().Interface().(web.UserSignUpRequest)

		if request.PasswordConfirm != request.Password {
			sl.ReportError(request.PasswordConfirm, "passwordConfirm", "PasswordConfirm", "passwordConfirm", "")
		}
	}, web.UserSignUpRequest{})
	service.Validate.AddTranslation("passwordConfirm", "Passwords do not match")

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	helper.PanicIfError(tx.Error)
	defer helper.CommitOrRollback(tx)

	return web.UserResponse{}
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
