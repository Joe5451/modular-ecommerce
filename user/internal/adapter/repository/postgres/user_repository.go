package postgres

import (
	"context"

	"github.com/Joe5451/modular-ecommerce/user/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/stackus/errors"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{db: db}
}

var _ domain.UserRepository = (*UserRepository)(nil)

func (r *UserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	query := `SELECT id, name, email, password FROM users WHERE id = $1`
	var user domain.User

	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.Wrap(errors.ErrNotFound, "user not found")
		}
		return nil, errors.Wrap(err, "failed to find user by id")
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	var user domain.User

	err := r.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // User not found is acceptable
		}
		return nil, errors.Wrap(err, "failed to find user by email")
	}

	return &user, nil
}

func (r *UserRepository) Save(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, query, user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return errors.Wrap(err, "failed to save user")
	}

	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	query := `UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4`
	_, err := r.db.Exec(ctx, query, user.Name, user.Email, user.Password, user.ID)
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	return nil
}
