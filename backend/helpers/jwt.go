package helpers

import (
	"bSocial/domain"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CreateJWTToken(user domain.User) (string, int64, error) {
	exp := time.Now().Add(time.Minute * time.Duration(CONFIG.Json.ExpMinute)).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = exp
	t, err := token.SignedString([]byte(CONFIG.Json.Secret))
	if err != nil {
		return "", 0, err
	}

	return t, exp, err
}

func ExtractJWTUserID(c *fiber.Ctx) uint {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["user_id"].(float64)
	return uint(id)
}
