package middleware

import (
	"fmt"
	"mfahmii/golang-restful/exception"
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
			fmt.Println("keluar dari if notFoundError")
			if validationErrors(ctx, r) {
				fmt.Println("di validationErrors")
				return
			}

			fmt.Println("keluar dari if validationErrors")
			if loginError(ctx, r) {
				return
			}
			fmt.Println("keluar dari if loginError")
			internalServerError(ctx, r)
			return
			// }
		}
	}()
	fmt.Println("lewat")
	return ctx.Next()
}

func validationErrors(ctx *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(exception.ValidationError)
	if ok {
		// ctx.Set("Content-Type", "application/json")
		ctx.Status(fiber.StatusBadRequest)

		webResponse := web.WebResponseError{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Errors: exception.Error(),
		}

		var _ = helper.WriteToResponseBody(ctx, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(ctx *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(exception.NotFoundError)
	if ok {
		// ctx.Set("Content-Type", "application/json")
		ctx.Status(fiber.StatusNotFound)

		webResponse := web.WebResponseError{
			Code:   fiber.StatusNotFound,
			Status: "NOT FOUND",
			Errors: exception.Error,
		}

		var _ = helper.WriteToResponseBody(ctx, webResponse)
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

	var _ = helper.WriteToResponseBody(ctx, webResponse)
}

func loginError(ctx *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(*exception.LoginError)
	if ok {
		// ctx.Set("Content-Type", "application/json")
		ctx.Status(fiber.StatusUnauthorized)

		webResponse := web.WebResponseError{
			Code:   fiber.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Errors: exception.Error(),
		}

		var _ = helper.WriteToResponseBody(ctx, webResponse)
		return true
	} else {
		return false
	}
}
