package repositories

import (
	"context"
	"leafall/todo-service/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

const MaxLimit = 100

type UserRepositoryInterface interface {
	FindAllPagination(ctx context.Context, limit, page int) ([]models.User, int64, error)
	FindByIdWithTasks(ctx context.Context, id int64, limit, page int) (*models.User, int64, error)
	FindById(ctx context.Context, id int64) (*models.User, error)
	Create(ctx context.Context, user *models.User) (error)
	Update(ctx context.Context, user *models.User) (error)
	DeleteById(ctx context.Context, id int64) (error)
	FindByLogin(ctx context.Context, login string) (*models.User, error)
}

type UserRepository struct {
	db *sqlx.DB
	TaskRepository TaskRepositoryInterface
}

func (r *UserRepository) FindAllPagination(ctx context.Context, limit, page int) ([]models.User, int64, error) {
	FixIfNotValidateLimitAndPage(&limit, &page)
	offset := (page - 1) * limit

	var users []models.User
	var count int64
	query := "SELECT * FROM users ORDER BY id ASC LIMIT $1 OFFSET $2";
	err := r.db.SelectContext(ctx, &users, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	query = "SELECT COUNT(*) FROM users";
	err = r.db.GetContext(ctx, &count, query)

	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func (r *UserRepository) FindByIdWithTasks(ctx context.Context, id int64, limit, page int) (*models.User, int64, error) {
	var user models.User
	query := "SELECT u.* FROM users u WHERE id=$1";
	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, 0, err
	}
	tasks, totalCount, err := r.TaskRepository.FindAllByUserId(ctx, user.Id, limit, page)
	if err != nil {
		return nil, 0, err
	}
	user.Tasks = tasks
	return &user, totalCount, nil
}

func (r *UserRepository) FindById(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	query := "SELECT u.* FROM users u WHERE id=$1";
	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByLogin(ctx context.Context, login string) (*models.User, error) {
	var user models.User
	query := "SELECT u.* FROM users u WHERE login=$1";
	err := r.db.GetContext(ctx, &user, query, login)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	user.CreatedAt = time.Now().UTC()
	query := "INSERT INTO users (login, password, created_at) VALUES ($1, $2, $3) RETURNING id";
	return r.db.GetContext(ctx, user, query, user.Login, user.Password, user.CreatedAt)
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) (error) {
	query := "UPDATE users SET login=$1, password=$2 WHERE id=$3";
	_, err := r.db.ExecContext(ctx, query, user.Login, user.Password, user.Id)
	return err
}

func (r *UserRepository) DeleteById(ctx context.Context, id int64) (error) {
	query := "DELETE FROM users WHERE id=$1";
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

