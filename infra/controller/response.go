package controller

import "github.com/gofiber/fiber"

func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status":  "Error",
		"data":    "",
		"message": err.Error(),
	}
}

// TODO: improve this success response
func SuccessResponse(data, message interface{}) *fiber.Map {
	return &fiber.Map{
		"status":  "Success",
		"data":    data,
		"message": message,
	}
}