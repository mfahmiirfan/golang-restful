package exception

import (
	"mfahmii/golang-restful/helper"
	"mfahmii/golang-restful/model/web"

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
	defer func() {
		if r := recover(); r != nil {
			// if err := ctx.Next(); err != nil {
			// Handle the error here
			if notFoundError(ctx, r) {
				return
			}

			if validationErrors(ctx, r) {
				return
			}

			if loginError(ctx, r) {
				return
			}

			internalServerError(ctx, r)
			// }
		}
	}()
	return ctx.Next()
}

func validationErrors(ctx *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(*ValidationError)
	if ok {
		// ctx.Set("Content-Type", "application/json")
		ctx.Status(fiber.StatusBadRequest)

		webResponse := web.WebResponseError{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Errors: exception.Error(),
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

		webResponse := web.WebResponseError{
			Code:   fiber.StatusNotFound,
			Status: "NOT FOUND",
			Errors: exception.Error,
		}

		helper.WriteToResponseBody(ctx, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(ctx *fiber.Ctx, err interface{}) {
	// ctx.Set("Content-Type", "application/json")
	// ctx.Status(fiber.StatusInternalServerError)

	// webResponse := web.WebResponse{
	// 	Code:   fiber.StatusInternalServerError,
	// 	Status: "INTERNAL SERVER ERROR",
	// 	Data:   err,
	// }

	// helper.WriteToResponseBody(ctx, webResponse)
	data := func() interface{} {
		if err, ok := err.(error); ok {
			return err.Error()
		}
		return err
	}()

	ctx.Status(fiber.StatusInternalServerError)

	webResponse := web.WebResponseError{
		Code:   fiber.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Errors: data,
	}

	helper.WriteToResponseBody(ctx, webResponse)
}

func loginError(ctx *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(LoginError)
	if ok {
		// ctx.Set("Content-Type", "application/json")
		ctx.Status(fiber.StatusUnauthorized)

		webResponse := web.WebResponseError{
			Code:   fiber.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Errors: exception.Error(),
		}

		helper.WriteToResponseBody(ctx, webResponse)
		return true
	} else {
		return false
	}
}
