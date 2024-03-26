package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sznborges/to_do_list/infra/controller"
)

func TaskRouter(app fiber.Router, controller controller.TaskController) {
	app.Get("/orders", controller.FindAll)
}
