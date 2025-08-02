package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetUserFromCtx(ctx *fiber.Ctx) (string, error) {
	user := ctx.Locals("userClaims")
	if user == nil {
		return "", errors.New("no user in context")
	}

	claims, ok := user.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid user claims type")
	}

	return claims["sub"].(string), nil
}
