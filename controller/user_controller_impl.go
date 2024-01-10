package controller

import (
	"mfahmii/golang-restful/helper"
	"mfahmii/golang-restful/model/web"
	"mfahmii/golang-restful/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(ctx *fiber.Ctx) error {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(ctx, &userCreateRequest)
	userResponse := controller.UserService.Create(ctx.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return helper.WriteToResponseBody(ctx, webResponse)
}

func (controller *UserControllerImpl) Update(ctx *fiber.Ctx) error {
	userUpdateRequest := web.UserUpdateRequest{}
	helper.ReadFromRequestBody(ctx, &userUpdateRequest)

	userId := ctx.Params("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.ID = id

	userResponse := controller.UserService.Update(ctx.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return helper.WriteToResponseBody(ctx, webResponse)
}

func (controller *UserControllerImpl) Delete(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.Delete(ctx.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	return helper.WriteToResponseBody(ctx, webResponse)
}

func (controller *UserControllerImpl) FindById(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindById(ctx.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return helper.WriteToResponseBody(ctx, webResponse)
}

func (controller *UserControllerImpl) FindAll(ctx *fiber.Ctx) error {
	userResponses := controller.UserService.FindAll(ctx.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponses,
	}

	return helper.WriteToResponseBody(ctx, webResponse)
}
