package repositories

import "github.com/jmoiron/sqlx"

func FixIfNotValidateLimitAndPage(limit, page *int) {
	if *limit <= 0 {
		*limit = 10
	}
	if *page <= 0 {
		*page = 1
	}
	if *limit > MaxLimit {
		*limit = MaxLimit
	}
}

func CreateTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func CreateUserRepository(db *sqlx.DB, taskRepository TaskRepositoryInterface) *UserRepository {
	return &UserRepository{db: db, TaskRepository: taskRepository}
}