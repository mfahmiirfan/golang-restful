package helper

import (
	"mfahmii/golang-restful/model/domain"
	"mfahmii/golang-restful/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		ID:       *user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     *user.Role,
		Verified: *user.Verified,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}
