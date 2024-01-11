package router

import (
	"mfahmii/golang-restful/app"
	"mfahmii/golang-restful/controller"
	"mfahmii/golang-restful/middleware"
	"mfahmii/golang-restful/service"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(categoryController controller.CategoryController,
	userController controller.UserController,
	authController controller.AuthController,
	config *app.Config,
	userService service.UserService) *fiber.App {
	// router := httprouter.New()
	router := fiber.New()

	// router.Use(middleware.ErrorHandler)
	router.Use(middleware.ErrorHandler)

	router.Get("/api/categories", categoryController.FindAll)
	router.Get("/api/categories/:categoryId", categoryController.FindById)
	router.Post("/api/categories", middleware.AuthHandler(config, userService), categoryController.Create)
	router.Put("/api/categories/:categoryId", middleware.AuthHandler(config, userService), categoryController.Update)
	router.Delete("/api/categories/:categoryId", middleware.AuthHandler(config, userService), categoryController.Delete)

	router.Get("/api/users", middleware.AuthHandler(config, userService), userController.FindAll)
	router.Get("/api/users/:userId", middleware.AuthHandler(config, userService), userController.FindById)
	router.Post("/api/users", middleware.AuthHandler(config, userService), userController.Create)
	router.Put("/api/users/:userId", middleware.AuthHandler(config, userService), userController.Update)
	router.Delete("/api/users/:userId", middleware.AuthHandler(config, userService), userController.Delete)

	router.Post("/api/auth/signup", authController.SignUp)
	router.Post("/api/auth/signin", authController.SignIn)
	router.Post("/api/auth/signout", middleware.AuthHandler(config, userService), authController.SignOut)

	// router.PanicHandler = exception.ErrorHandler

	return router
}
