package helper

import (
	"github.com/gofiber/fiber/v2"
)

func ReadFromRequestBody(ctx *fiber.Ctx, result interface{}) {
	// decoder := json.NewDecoder(request.Body)
	// err := decoder.Decode(result)
	err := ctx.BodyParser(result)
	PanicIfError(err)
}

func WriteToResponseBody(ctx *fiber.Ctx, response interface{}) error {
	ctx.Append("Content-Type", "application/json")
	// encoder := json.NewEncoder(writer)
	// err := encoder.Encode(response)
	err := ctx.JSON(response)
	PanicIfError(err)
	return err
}
