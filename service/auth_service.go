package service

import (
	"mfahmii/golang-restful/model/web"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	SignUp(ctx *fiber.Ctx, request web.UserSignUpRequest) web.UserResponse
	SignIn(ctx *fiber.Ctx, request web.UserSignInRequest) web.TokenResponse
}
