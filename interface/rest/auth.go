package rest

import (
	"bSocial/domain"
	"bSocial/helpers"
	"bSocial/interface/mysql"

	"github.com/gofiber/fiber/v2"
)

func InitAuthApi(api fiber.Router) {
	api.Post("/register", apiRegister())
	api.Post("/login", apiLogin())
}

func apiRegister() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userData := new(domain.ApiRegister)
		if err := c.BodyParser(userData); err != nil {
			return err
		}
		if errors := domain.ValidateType(*userData); errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)
		}
		passHash, _ := helpers.HashPassword(userData.ReqPassword)
		userData.Password = passHash
		newUser, _ := mysql.CreateUser(userData.User)

		return c.JSON(newUser)
	}
}

func apiLogin() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		loginData := new(domain.ApiLogin)
		if err := c.BodyParser(loginData); err != nil {
			return err
		}
		if errors := domain.ValidateType(*loginData); errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)
		}

		user, err := mysql.GetUserByUsernameAndPassword(loginData.Username, loginData.Password)
		if err != nil {
			return err
		}
		token, exp, err := helpers.CreateJWTToken(*user)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})
	}
}
