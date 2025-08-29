package repositories

import (
	"context"
	"leafall/todo-service/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type TaskRepositoryInterface interface {
	FindAllByUserId(ctx context.Context, userId int64, limit, page int) ([]models.Task, int64, error)
	FindById(ctx context.Context, id int64) (*models.Task, error)
	Create(ctx context.Context, task *models.Task) (error)
	Update(ctx context.Context, task *models.Task) (error)
	DeleteById(ctx context.Context, id int64) (error)
}

type TaskRepository struct {
	db *sqlx.DB
}

func (r *TaskRepository) FindAllByUserId(ctx context.Context, userId int64, limit, page int) ([]models.Task, int64, error) {
	FixIfNotValidateLimitAndPage(&limit, &page)
	var tasks []models.Task
	var totalCount int64
	query := "SELECT * FROM tasks WHERE userId = $1 LIMIT $2 OFFSET $3";
	err := r.db.SelectContext(ctx, &tasks, query, userId, limit, limit * page)
	if err != nil {
		return nil, 0, err
	}
	query = "SELECT COUNT(*) FROM tasks WHERE userId = $1";
	err = r.db.GetContext(ctx, &totalCount, query, userId)
	if err != nil {
		return nil, 0, err
	}
	return tasks, totalCount, nil
}

func (r *TaskRepository) FindById(ctx context.Context, id int64) (*models.Task, error) {
	var task models.Task
	query := "SELECT * FROM tasks WHERE id=$1";
	err := r.db.GetContext(ctx, &task, query, id)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Create(ctx context.Context, task *models.Task) (error) {
	task.CreatedAt = time.Now().UTC()
	query := "INSERT INTO tasks (name, description, created_at, user_id) VALUES ($1, $2, $3, $4)";
	return r.db.GetContext(ctx, task, query, task.Name, task.Description, task.CreatedAt, task.UserId)
}

func (r *TaskRepository) Update(ctx context.Context, task *models.Task) (error) {
	task.UpdatedAt = time.Now().UTC()
	query := "UPDATE tasks SET name=$1, description=$2, updated_at=$3 WHERE id=$4";
	_, err := r.db.ExecContext(ctx, query, task.Name, task.Description, task.UpdatedAt, task.Id)
	return err;
}

func (r *TaskRepository) DeleteById(ctx context.Context, id int64) (error) {
	query := "DELETE FROM tasks WHERE id=$1";
	_, err := r.db.ExecContext(ctx, query, id)
	return err;
}

