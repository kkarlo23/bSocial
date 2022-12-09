package rest

import (
	"bSocial/helpers"
	"bSocial/interface/mysql"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func InitUserApi(api fiber.Router) {
	api.Get("/user", apiGetUsers())
	api.Post("/user/follow/:userID", apiPostUserFollow())
}

func apiGetUsers() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		users, err := mysql.GetUsers()
		if err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}
		return ResponseWithData(c, users)
	}
}

func apiPostUserFollow() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		thisUserID := helpers.ExtractJWTUserID(c)
		getUserIDInt64, err := strconv.ParseInt(c.Params("userID"), 10, 32)
		if err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}
		getUserID := uint(getUserIDInt64)

		err = mysql.UserFollow(thisUserID, getUserID)
		if err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}
		return ResponseWithData(c, 1)
	}
}
