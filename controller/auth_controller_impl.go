package controller

import (
	"mfahmii/golang-restful/helper"
	"mfahmii/golang-restful/model/web"
	"mfahmii/golang-restful/service"

	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) SignUp(ctx *fiber.Ctx) error {
	// userSignUpRequest := web.UserSignUpRequest{}
	// helper.ReadFromRequestBody(ctx, &userSignUpRequest)

	// if userSignUpRequest.Password != userSignUpRequest.PasswordConfirm {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwords do not match"})

	// }

	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userSignUpRequest.Password), bcrypt.DefaultCost)
	// helper.PanicIfError(err)

	// // if err != nil {
	// // 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	// // }
	// newUser := web.UserCreateRequest{
	// 	Name:     userSignUpRequest.Name,
	// 	Email:    strings.ToLower(userSignUpRequest.Email),
	// 	Password: string(hashedPassword),
	// 	// Photo:    &payload.Photo,
	// }

	// userResponse := controller.AuthService.Create(ctx.Context(), newUser)

	// // categoryResponse := controller.AuthService.Create(ctx.Context(), userSignUpRequest)
	// webResponse := web.WebResponse{
	// 	Code:   200,
	// 	Status: "OK",
	// 	Data:   userResponse,
	// }

	// return helper.WriteToResponseBody(ctx, webResponse)
	return nil
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
	userResponse := controller.AuthService.SignIn(ctx, userSignInRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return helper.WriteToResponseBody(ctx, webResponse)
}

func (controller *AuthControllerImpl) Logout(ctx *fiber.Ctx) error {
	// categoryId := ctx.Params("categoryId")
	// id, err := strconv.Atoi(categoryId)
	// helper.PanicIfError(err)

	// controller.AuthService.Delete(ctx.Context(), id)
	// webResponse := web.WebResponse{
	// 	Code:   200,
	// 	Status: "OK",
	// }

	// return helper.WriteToResponseBody(ctx, webResponse)
	return nil
}
