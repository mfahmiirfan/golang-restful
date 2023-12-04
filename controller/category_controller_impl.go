package controller

import (
	"mfahmii/golang-restful/helper"
	"mfahmii/golang-restful/model/web"
	"mfahmii/golang-restful/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(ctx *fiber.Ctx) error {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(ctx, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(ctx.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	return helper.WriteToResponseBody(ctx, webResponse)
}

func (controller *CategoryControllerImpl) Update(ctx *fiber.Ctx) error {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(ctx, &categoryUpdateRequest)

	categoryId := ctx.Params("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(ctx.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	return helper.WriteToResponseBody(ctx, webResponse)
}

func (controller *CategoryControllerImpl) Delete(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(ctx.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	return helper.WriteToResponseBody(ctx, webResponse)
}

func (controller *CategoryControllerImpl) FindById(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(ctx.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	return helper.WriteToResponseBody(ctx, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(ctx *fiber.Ctx) error {
	categoryResponses := controller.CategoryService.FindAll(ctx.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	return helper.WriteToResponseBody(ctx, webResponse)
}
