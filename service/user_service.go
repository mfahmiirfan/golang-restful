package service

import (
	"context"
	"mfahmii/golang-restful/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) web.UserResponse
	FindByEmail(ctx context.Context, email string) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}
