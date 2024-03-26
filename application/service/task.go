package service

import (
	"github.com/sznborges/to_do_list/domain"
	"github.com/sznborges/to_do_list/domain/dto"
)

type Task struct {
	repository domain.FindTaskRepository
}

func NewTask(repository domain.FindTaskRepository) *Task {
	return &Task{repository: repository}
}

func (s *Task) FindAll(input dto.InputTaskDto) (*dto.OutputTaskDto, error) {
	tasks, err := s.repository.FindAll(input)
	var taskList []dto.TaskDto
	if err != nil {
		return nil, err
	}
	if len(tasks.Tasks) > 0 {
		for _, tsk := range tasks.Tasks {
			temp := dto.TaskDto{
				ID:          tsk.ID,
				Title:       tsk.Title,
				Description: tsk.Description,
				Completed:   tsk.Completed,
				CreatedAt:   tsk.CreatedAt,
			}
			taskList = append(taskList, temp)
		}
	}
	output := dto.OutputTaskDto{Tasks: taskList}
	return &output, nil
}
