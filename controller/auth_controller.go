package controller

import (
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	SignUp(ctx *fiber.Ctx) error
	SignIn(ctx *fiber.Ctx) error
	SignOut(ctx *fiber.Ctx) error
}
