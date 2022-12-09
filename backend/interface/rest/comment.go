package rest

import (
	"bSocial/domain"
	"bSocial/interface/mysql"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func InitCommentApi(api fiber.Router) {
	api.Post("/comment/:postID", apiPostComment())
}

// creates a user comment for specific postID
func apiPostComment() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		getIntPostID, err := strconv.ParseInt(c.Params("postID"), 10, 32)
		if err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}
		commentData := new(domain.Comment)
		if err := c.BodyParser(commentData); err != nil {
			return ResponseWithError(c, err.Error(), nil)
		}
		if errors := domain.ValidateType(*commentData); errors != nil {
			return ResponseWithError(c, "", errors)
		}
		commentData.PostID = uint(getIntPostID)
		newComment, _ := mysql.CreateComment(commentData)

		return ResponseWithData(c, newComment)
	}
}
