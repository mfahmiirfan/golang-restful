package service

import (
	"context"
	"mfahmii/golang-restful/model/web"
)

type AuthService interface {
	FindUserByEmail(ctx context.Context, email string) web.UserResponse
}
