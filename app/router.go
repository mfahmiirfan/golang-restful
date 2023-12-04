package app

import (
	"mfahmii/golang-restful/controller"
	"mfahmii/golang-restful/exception"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(categoryController controller.CategoryController) *fiber.App {
	// router := httprouter.New()
	router := fiber.New()

	router.Get("/api/categories", categoryController.FindAll)
	router.Get("/api/categories/:categoryId", categoryController.FindById)
	router.Post("/api/categories", categoryController.Create)
	router.Put("/api/categories/:categoryId", categoryController.Update)
	router.Delete("/api/categories/:categoryId", categoryController.Delete)

	// router.PanicHandler = exception.ErrorHandler
	router.Use(exception.ErrorHandler)

	return router
}
