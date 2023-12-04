package service

import (
	"context"
	"mfahmii/golang-restful/model/web"

	"github.com/google/uuid"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, userId uuid.UUID)
	FindById(ctx context.Context, userId uuid.UUID) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}
