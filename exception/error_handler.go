package exception

import (
	"mfahmii/golang-restful/helper"
	"mfahmii/golang-restful/model/web"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

// 	if notFoundError(writer, request, err) {
// 		return
// 	}

// 	if validationErrors(writer, request, err) {
// 		return
// 	}

//		internalServerError(writer, request, err)
//	}
func ErrorHandler(ctx *fiber.Ctx) error {
	if err := ctx.Next(); err != nil {
		// Handle the error here
		if notFoundError(ctx, err) {
			return err
		}

		if validationErrors(ctx, err) {
			return err
		}

		internalServerError(ctx, err)
		return err
	}
	return nil
}

func validationErrors(ctx *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		// ctx.Set("Content-Type", "application/json")
		ctx.Status(fiber.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(ctx, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(ctx *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		// ctx.Set("Content-Type", "application/json")
		ctx.Status(fiber.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   fiber.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(ctx, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(ctx *fiber.Ctx, err interface{}) {
	// ctx.Set("Content-Type", "application/json")
	ctx.Status(fiber.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   fiber.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(ctx, webResponse)
}
