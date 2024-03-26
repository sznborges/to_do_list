package dto

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sznborges/to_do_list/infra/logger"
)

func WriteResponse(c *fiber.Ctx, body any, status int) error {
	if body == nil {
		return c.SendStatus(status)
	}
	bytes, err := json.Marshal(body)
	if err != nil {
		logger.Logger.WithError(err).Error("error marshaling json body")
		return c.SendStatus(http.StatusInternalServerError)
	}
	c.Context().Response.SetBodyRaw(bytes)
	c.Context().Response.Header.SetContentType("application/json")
	return c.SendStatus(status)
}

// WriteError to fiber context
func WriteError(c *fiber.Ctx, err error) error {
	logger.Logger.WithError(err).Error("unknown error")
	return WriteResponse(c, map[string]any{
		"Code":    "001",
		"Message": err.Error(),
	}, http.StatusInternalServerError)
}