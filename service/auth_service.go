package service

import (
	"context"
	"mfahmii/golang-restful/model/web"
)

type AuthService interface {
	SignUp(ctx context.Context, request web.UserSignUpRequest) web.UserResponse
	SignIn(ctx context.Context, request web.UserSignInRequest) web.TokenResponse
}
