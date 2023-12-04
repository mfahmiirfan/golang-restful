package main

import (
	"mfahmii/golang-restful/app"
	"mfahmii/golang-restful/controller"
	"mfahmii/golang-restful/helper"
	"mfahmii/golang-restful/repository"
	"mfahmii/golang-restful/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// db := app.NewDB()
	db := app.OpenConnection()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	// server := http.Server{
	// 	Addr:    "localhost:3000",
	// 	Handler: middleware.NewAuthMiddleware(router),
	// }
	err := router.Listen(":3000")

	// err := server.ListenAndServe()
	helper.PanicIfError(err)
}
