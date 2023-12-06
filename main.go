package main

import (
	"mfahmii/golang-restful/app"
	"mfahmii/golang-restful/config"
	"mfahmii/golang-restful/controller"
	"mfahmii/golang-restful/helper"
	"mfahmii/golang-restful/model/domain"
	"mfahmii/golang-restful/repository"
	"mfahmii/golang-restful/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := config.LoadConfig("./config")
	helper.PanicIfError(err)
	// db := app.NewDB()
	db := app.OpenConnection()
	db.AutoMigrate(&domain.User{})
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	userRepository := repository.NewUserRepository()
	authService := service.NewAuthService(userRepository, db, validate, &config)
	authController := controller.NewAuthController(authService)
	router := app.NewRouter(categoryController, authController)

	// server := http.Server{
	// 	Addr:    "localhost:3000",
	// 	Handler: middleware.NewAuthMiddleware(router),
	// }
	err = router.Listen(":3000")

	// err := server.ListenAndServe()
	helper.PanicIfError(err)
}
