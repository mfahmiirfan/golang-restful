package middleware

import (
	"fmt"
	"mfahmii/golang-restful/app"
	"mfahmii/golang-restful/exception"
	"mfahmii/golang-restful/service"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AuthHandler(config *app.Config, userService service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("middleware auth")
		var tokenString string
		authorization := c.Get("Authorization")

		if strings.HasPrefix(authorization, "Bearer ") {
			tokenString = strings.TrimPrefix(authorization, "Bearer ")
		} else if c.Cookies("token") != "" {
			tokenString = c.Cookies("token")
		}

		if tokenString == "" {
			// return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
			panic(exception.NewUnauthorizeError("You are not logged in"))
		}

		tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
			}

			return []byte(config.JwtSecret), nil
		})
		if err != nil {
			// return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
			panic(exception.NewUnauthorizeError(fmt.Sprintf("invalidate token: %v", err)))
		}

		claims, ok := tokenByte.Claims.(jwt.MapClaims)
		if !ok || !tokenByte.Valid {
			// return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})
			panic(exception.NewUnauthorizeError("invalid token claim"))
		}

		fmt.Println(">")
		fmt.Println(claims["role"])

		// id, err := strconv.Atoi(claims["sub"])
		// helper.PanicIfError(err)
		executedOnce := false

		defer func() {
			fmt.Println("masuk defer")
			if r := recover(); r != nil {
				fmt.Println("masuk recover")
				if _, ok := r.(exception.NotFoundError); ok && !executedOnce {
					panic(exception.NewUnauthorizeError("the user belonging to this token no longer exists"))
				}
				fmt.Println(reflect.TypeOf(r))
				panic(r)
			}
		}()
		fmt.Println("masuk pak eko")

		userResponse := userService.FindById(c.Context(), int(claims["sub"].(float64)))
		fmt.Println("coba di tes")
		executedOnce = true

		if checkPermission(claims["role"].(string), c.Path(), c.Method()) {
			c.Locals("user", userResponse)

			return c.Next()
		}

		// if strconv.Itoa(*user.ID) != claims["sub"] {
		// 	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
		// }
		fmt.Println("keluar")
		panic(exception.NewForbiddenError("Access to the requested resource is forbidden."))
	}
}

func checkPermission(role, endpoint, method string) bool {
	fmt.Println(role)
	if strings.HasPrefix(endpoint, "/api/categories") && role == "admin" {
		return true
	}

	if strings.HasPrefix(endpoint, "/api/users") && role == "admin" {
		return true
	}

	if strings.HasPrefix(endpoint, "/api/auth/signout") {
		return true
	}

	return false
}
