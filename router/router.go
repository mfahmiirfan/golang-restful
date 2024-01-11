package router

import (
	"mfahmii/golang-restful/app"
	"mfahmii/golang-restful/controller"
	"mfahmii/golang-restful/middleware"
	"mfahmii/golang-restful/service"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(categoryController controller.CategoryController, authController controller.AuthController, config *app.Config, userService service.UserService) *fiber.App {
	// router := httprouter.New()
	router := fiber.New()

	// router.Use(middleware.ErrorHandler)
	router.Use(middleware.ErrorHandler, middleware.AuthHandler(config, userService))

	router.Get("/api/categories", categoryController.FindAll)
	router.Get("/api/categories/:categoryId" /*, middleware.AuthHandler(config, userService)*/, categoryController.FindById)
	router.Post("/api/categories", categoryController.Create)
	router.Put("/api/categories/:categoryId", categoryController.Update)
	router.Delete("/api/categories/:categoryId", categoryController.Delete)
	router.Post("/api/auth/signup", authController.SignUp)
	router.Post("/api/auth/signin", authController.SignIn)
	router.Post("/api/auth/signout", authController.SignOut)

	// router.PanicHandler = exception.ErrorHandler

	return router
}
