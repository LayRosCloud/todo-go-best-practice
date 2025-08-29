package repositories

import (
	"context"
	"leafall/todo-service/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type TokenRepositoryInterface interface {
	Create(ctx context.Context, token *models.Token, duration time.Duration) error
	FindByToken(ctx context.Context, token string) (*models.Token, error)
	DeleteToken(ctx context.Context, token string) error
}

type TokenRepository struct {
	db *sqlx.DB
}

func (t *TokenRepository) Create(ctx context.Context, token *models.Token, duration time.Duration) error {
	token.CreatedAt = time.Now()
	token.ExpiredAt = time.Now().Add(duration)
	query := "INSERT INTO tokens(token, user_id, created_at, expired_at) VALUES ($1, $2, $3, $4) RETURNING id"
	return t.db.GetContext(ctx, token, query, token.Token, token.UserId, token.CreatedAt, token.ExpiredAt)
}

func (t *TokenRepository) FindByToken(ctx context.Context, token string) (*models.Token, error) {
	var entity *models.Token
	query := "SELECT * FROM tokens WHERE token=$1"
	err := t.db.GetContext(ctx, entity, query, token)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (t *TokenRepository) DeleteToken(ctx context.Context, token string) error {
	query := "DELETE FROM tokens WHERE token=$1"
	_, err := t.db.ExecContext(ctx, query, token)
	return err
}