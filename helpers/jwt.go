package helpers

import (
	"bSocial/domain"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const EXP_IN_MINUTES = 30

func CreateJWTToken(user domain.User) (string, int64, error) {
	exp := time.Now().Add(time.Minute * EXP_IN_MINUTES).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = exp
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}
	return t, exp, err
}
