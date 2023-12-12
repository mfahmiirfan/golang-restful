package main

import (
	"mfahmii/golang-restful/app"
	"mfahmii/golang-restful/controller"
	"mfahmii/golang-restful/helper"
	"mfahmii/golang-restful/repository"
	"mfahmii/golang-restful/router"
	"mfahmii/golang-restful/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := app.NewConfig(".")
	// db := app.NewDB()
	db := app.NewDB(config)

	validation := app.NewValidation()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validation)
	categoryController := controller.NewCategoryController(categoryService)

	userRepository := repository.NewUserRepository()
	authService := service.NewAuthService(userRepository, db, validation, config)
	authController := controller.NewAuthController(authService)

	router := router.NewRouter(categoryController, authController)

	// server := http.Server{
	// 	Addr:    "localhost:3000",
	// 	Handler: middleware.NewAuthMiddleware(router),
	// }
	err := router.Listen(":3000")

	// err := server.ListenAndServe()
	helper.PanicIfError(err)
}
