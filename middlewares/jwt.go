package middlewares

import (
	"os"
	"testku/helpers"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return helpers.ErrorResponse(c, "Unauthorized", fiber.StatusUnauthorized, nil)
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			// Extract the token from the context
			userToken := ctx.Locals("user").(*jwt.Token)
			claims := userToken.Claims.(jwt.MapClaims)

			// Extract the user ID from the token
			ctx.Locals("email", claims["email"])
			ctx.Locals("role", claims["role"])
			return ctx.Next()
		},
	})
}
