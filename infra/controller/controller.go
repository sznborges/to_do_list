package controller

import "github.com/gofiber/fiber/v2"

type BaseController interface {
	FindAll(ctx *fiber.Ctx) error
}

type Controller interface {
	FindAll(ctx *fiber.Ctx) error
}
