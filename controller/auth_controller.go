package controller

import (
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	SignUp(ctx *fiber.Ctx) error
	SignIn(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
}
