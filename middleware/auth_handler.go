package middleware

import (
	"fmt"
	"mfahmii/golang-restful/app"
	"mfahmii/golang-restful/service"
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
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
		}

		tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
			}

			return []byte(config.JwtSecret), nil
		})
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
		}

		claims, ok := tokenByte.Claims.(jwt.MapClaims)
		if !ok || !tokenByte.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})

		}

		fmt.Println(claims["sub"])

		// id, err := strconv.Atoi(claims["sub"])
		// helper.PanicIfError(err)

		userResponse := userService.FindById(c.Context(), claims["sub"].(int))

		// if strconv.Itoa(*user.ID) != claims["sub"] {
		// 	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
		// }

		c.Locals("user", userResponse)

		return c.Next()
	}
}
