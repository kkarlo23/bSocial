package rest

import (
	"bSocial/domain"

	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Message    string                  `json:"message"`
	Validation []*domain.ErrorValidate `json:"validation"`
}

// We could also go without Success field, since it is kinda redundant because Error can be nil
type Response struct {
	Success bool           `json:"success"`
	Error   *ErrorResponse `json:"error"`
	Data    interface{}    `json:"data"`
}

// RespopnseWithError standardizes rest Success responses, so it have similar structure troughout API
// It is better to use more variety of statuses, they describe what happened in general
// But for this case we are going to use Status OK (200) for everything
func ResponseWithData(c *fiber.Ctx, data interface{}) error {
	response := Response{Data: data, Success: true, Error: nil}
	return c.Status(fiber.StatusOK).JSON(response)
}

// RespopnseWithError standardizes rest Error responses, so it have similar structure troughout API
// It is better to use more variety of statuses, they describe what happened in general
// But for this case we are going to use Status Bad Request (400) for everything
func ResponseWithError(c *fiber.Ctx, message string, validation []*domain.ErrorValidate) error {
	if validation == nil {
		response := Response{Data: nil, Success: false, Error: &ErrorResponse{Message: message, Validation: validation}}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	response := Response{Data: nil, Success: false, Error: &ErrorResponse{Message: "Validation Error", Validation: validation}}
	return c.Status(fiber.StatusBadRequest).JSON(response)
}
