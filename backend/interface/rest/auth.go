package rest

import (
	"bSocial/domain"
	"bSocial/helpers"
	"bSocial/interface/kafkaProducer"
	"bSocial/interface/mysql"

	"github.com/gofiber/fiber/v2"
)

func InitAuthApi(api fiber.Router) {
	api.Post("/register", apiRegister())
	api.Post("/login", apiLogin())
}

// creates a new user
func apiRegister() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userData := new(domain.ApiRegister)
		if err := c.BodyParser(userData); err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}
		if errors := domain.ValidateType(*userData); errors != nil {
			return ResponseWithError(c, "", errors)
		}
		passHash, err := helpers.HashPassword(userData.ReqPassword)
		if err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}
		userData.Password = passHash
		newUser, err := mysql.CreateUser(userData.User)
		if err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}

		err = mysql.UserFollow(newUser.ID, newUser.ID)
		if err != nil {
			return err
		}
		// TODO: handle error
		kafkaProducer.ProduceUserRegister(*newUser)

		return ResponseWithData(c, newUser)
	}
}

// returns jwt if user data is valid
func apiLogin() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		loginData := new(domain.ApiLogin)
		if err := c.BodyParser(loginData); err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}
		if errors := domain.ValidateType(*loginData); errors != nil {
			return ResponseWithError(c, "", errors)
		}

		user, err := mysql.GetUserByUsernameAndPassword(loginData.Username, loginData.Password)
		if err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}
		token, exp, err := helpers.CreateJWTToken(*user)
		if err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}
		return ResponseWithData(c, fiber.Map{"token": token, "exp": exp, "user": user})
	}
}
