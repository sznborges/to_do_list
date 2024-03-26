package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sznborges/to_do_list/application/service"
	"github.com/sznborges/to_do_list/domain/dto"
)

type TaskController struct {
	svc *service.Task
}

func NewTaskController(svc *service.Task) *TaskController {
	return &TaskController{svc: svc}
}

// @Summary	Returns paged orders
// @Description Returns paged orders
// @Tags Order
// @Accept json
// @Produce json
// @Param ordersCode query string false "Codes"
// @Param page query int true "Page"
// @success 200
// @Router /api/v1/orders [get]
// @Security AuthKey
func (c *TaskController) FindAll(ctx *fiber.Ctx) error {
	input := dto.InputTaskDto{
		Tasks: dto.TaskDto{
			ID:          0,
			Title:       "",
			Description: "",
			Completed:   false,
			CreatedAt:   "",
		},
	}
	output, err := c.svc.FindAll(input)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(ErrorResponse(err))
	}
	return ctx.Status(http.StatusOK).JSON(SuccessResponse(output, nil))
}