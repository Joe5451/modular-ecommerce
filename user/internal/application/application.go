package application

import (
	"context"

	"github.com/Joe5451/modular-ecommerce/user/internal/domain"
)

type (
	RegisterUser struct {
		ID       string
		Name     string
		Email    string
		Password string
	}

	AuthenticateUser struct {
		Email    string
		Password string
	}

	GetUser struct {
		ID string
	}

	App interface {
		RegisterUser(ctx context.Context, cmd RegisterUser) error
		AuthenticateUser(ctx context.Context, cmd AuthenticateUser) (*domain.User, error)
		GetUser(ctx context.Context, query GetUser) (*domain.User, error)
	}

	Application struct {
		repo domain.UserRepository
	}
)

var _ App = (*Application)(nil)

func New(repo domain.UserRepository) *Application {
	return &Application{
		repo: repo,
	}
}

func (a *Application) RegisterUser(ctx context.Context, cmd RegisterUser) error {
	existingUser, err := a.repo.FindByEmail(ctx, cmd.Email)
	if err == nil && existingUser != nil {
		return domain.ErrDuplicateEmail
	}

	user, err := domain.RegisterUser(cmd.ID, cmd.Name, cmd.Email, cmd.Password)
	if err != nil {
		return err
	}

	return a.repo.Save(ctx, user)
}

func (a *Application) AuthenticateUser(ctx context.Context, cmd AuthenticateUser) (*domain.User, error) {
	user, err := a.repo.FindByEmail(ctx, cmd.Email)
	if err != nil {
		return nil, err
	}

	if err := user.Authenticate(cmd.Password); err != nil {
		return nil, err
	}

	return user, nil
}

func (a *Application) GetUser(ctx context.Context, query GetUser) (*domain.User, error) {
	return a.repo.FindByID(ctx, query.ID)
}
