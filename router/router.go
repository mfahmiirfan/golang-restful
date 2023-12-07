package router

import (
	"mfahmii/golang-restful/controller"
	"mfahmii/golang-restful/exception"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(categoryController controller.CategoryController, authController controller.AuthController) *fiber.App {
	// router := httprouter.New()
	router := fiber.New()

	router.Use(exception.ErrorHandler)

	router.Get("/api/categories", categoryController.FindAll)
	router.Get("/api/categories/:categoryId", categoryController.FindById)
	router.Post("/api/categories", categoryController.Create)
	router.Put("/api/categories/:categoryId", categoryController.Update)
	router.Delete("/api/categories/:categoryId", categoryController.Delete)
	router.Post("/api/auth/signup", authController.SignUp)
	router.Post("/api/auth/signin", authController.SignIn)
	router.Post("/api/auth/logout", authController.Logout)

	// router.PanicHandler = exception.ErrorHandler

	return router
}
