package controller

import (
	"mfahmii/golang-restful/app"
	"mfahmii/golang-restful/helper"
	"mfahmii/golang-restful/model/web"
	"mfahmii/golang-restful/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
	Config      *app.Config
}

func NewAuthController(authService service.AuthService, config *app.Config) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
		Config:      config,
	}
}

func (controller *AuthControllerImpl) SignUp(ctx *fiber.Ctx) error {
	userSignUpRequest := web.UserSignUpRequest{}
	helper.ReadFromRequestBody(ctx, &userSignUpRequest)

	userResponse := controller.AuthService.SignUp(ctx.Context(), userSignUpRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return helper.WriteToResponseBody(ctx, webResponse)
}

func (controller *AuthControllerImpl) SignIn(ctx *fiber.Ctx) (err error) {
	userSignInRequest := web.UserSignInRequest{}
	helper.ReadFromRequestBody(ctx, &userSignInRequest)

	// categoryId := ctx.Params("categoryId")
	// id, err := strconv.Atoi(categoryId)
	// helper.PanicIfError(err)

	// categoryUpdateRequest.Id = id

	// categoryResponse := controller.AuthService.Update(ctx.Context(), categoryUpdateRequest)
	// defer func() {
	// 	if r := recover(); r != nil {

	// 		if _, ok := r.(exception.NotFoundError); ok {
	// 			// Jika berhasil mengekstrak nilai error dari panic
	// 			panic(exception.NewLoginError("Invalid email or Password"))
	// 			// Lakukan sesuatu dengan nilai error
	// 		}
	// 		return
	// 	}
	// }()
	tokenResponse := controller.AuthService.SignIn(ctx.Context(), userSignInRequest)

	ctx.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenResponse.Token,
		Path:     "/",
		MaxAge:   controller.Config.JwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   tokenResponse,
	}

	return helper.WriteToResponseBody(ctx, webResponse)
}

func (controller *AuthControllerImpl) Logout(ctx *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	ctx.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	return helper.WriteToResponseBody(ctx, webResponse)
}
