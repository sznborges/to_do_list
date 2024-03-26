package repository

import (
	"database/sql"
	"time"

	"github.com/sznborges/to_do_list/domain"
	"github.com/sznborges/to_do_list/domain/dto"
	"github.com/sznborges/to_do_list/domain/entity"
)

type TaskRepository struct {
	conn *sql.DB
}

func NewTaskRepository(connector domain.DatabaseConnector) *TaskRepository {
	return &TaskRepository{conn: connector.GetConnection()}
}

func (r *TaskRepository) Create(task *entity.Task) error {
	query := `INSERT INTO tasks (id, title, description, completed, createdAt) VALUES ($1,$2,$3,$4,$5)`
	stmt, err := r.conn.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(task.ID, task.Title, task.Description, task.Completed, time.Now().String())
	if err != nil {
		return err
	}
	return err
}


// FindAll retrieves and returns a list of all orders for the order list page.
//
// Parameters:
//   - InputOrderListPageDto
//
// Returns:
//   - An array containing order objects.
//   - A map[string]int with total and lastPage.

func (r *TaskRepository) FindAll(input dto.InputTaskDto) (dto.OutputTaskDto,error) {
	var tasks []dto.TaskDto
	var queryParams []interface{}
	var err error
	baseQuery := `SELECT id, title, description, completed, createdAt count(1) OVER() AS full_count FROM tasks`
	whereClause := " WHERE 1=1 "
	orderBy := " ORDER BY created_at "
	finalQuery := baseQuery + whereClause + orderBy
	stmt, err := r.conn.Prepare(finalQuery)
	if err != nil {
		return dto.OutputTaskDto{Tasks: nil}, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(queryParams...)
	if err != nil {
		return dto.OutputTaskDto{Tasks: nil}, err
	}
	defer rows.Close()
	for rows.Next() {
		var taskObj dto.TaskDto
		err = rows.Scan(
			&taskObj.ID,
			&taskObj.Title,
			&taskObj.Description,
			&taskObj.Completed,
			&taskObj.CreatedAt)
		if err != nil {
			continue
		}
		tasks = append(tasks, taskObj)
	}
	output := dto.OutputTaskDto{Tasks: tasks}
	return output, nil
}