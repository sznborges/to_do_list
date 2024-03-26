package domain

import (
	"github.com/sznborges/to_do_list/domain/dto"
	"github.com/sznborges/to_do_list/domain/entity"
)

type CreateTaskRepository interface {
	Create(task *entity.Task) error
}

type FindTaskRepository interface {
	FindAll(input dto.InputTaskDto) (dto.OutputTaskDto, error)
}

type TaskRepository interface {
	CreateTaskRepository
	FindTaskRepository
}
