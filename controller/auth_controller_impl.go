package controller

import (
	"mfahmii/golang-restful/helper"
	"mfahmii/golang-restful/model/web"
	"mfahmii/golang-restful/service"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthControllerImpl struct {
	UserService service.UserService
}

func NewAuthController(categoryService service.UserService) AuthController {
	return &AuthControllerImpl{
		UserService: categoryService,
	}
}

func (controller *AuthControllerImpl) SignUp(ctx *fiber.Ctx) error {
	userSignUpRequest := web.UserSignUpRequest{}
	helper.ReadFromRequestBody(ctx, &userSignUpRequest)

	if userSignUpRequest.Password != userSignUpRequest.PasswordConfirm {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwords do not match"})

	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userSignUpRequest.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	// if err != nil {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	// }
	newUser := web.UserCreateRequest{
		Name:     userSignUpRequest.Name,
		Email:    strings.ToLower(userSignUpRequest.Email),
		Password: string(hashedPassword),
		// Photo:    &payload.Photo,
	}

	userResponse := controller.UserService.Create(ctx.Context(), newUser)

	// categoryResponse := controller.UserService.Create(ctx.Context(), userSignUpRequest)
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

	// categoryResponse := controller.UserService.Update(ctx.Context(), categoryUpdateRequest)
	defer func() {
		if r := recover(); r != nil {
			// fmt.Errorf("Recovered from panic: %v", r)
			webResponse := web.WebResponse{
				Code:   fiber.StatusNotFound,
				Status: "NOT FOUND",
				Data:   r,
			}
			err = helper.WriteToResponseBody(ctx, webResponse)
		}
	}()
	userResponse := controller.UserService.FindByEmail(ctx.Context(), userSignInRequest.Email)

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

	// controller.UserService.Delete(ctx.Context(), id)
	// webResponse := web.WebResponse{
	// 	Code:   200,
	// 	Status: "OK",
	// }

	// return helper.WriteToResponseBody(ctx, webResponse)
	return nil
}
