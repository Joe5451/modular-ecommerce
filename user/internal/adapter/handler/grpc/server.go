package grpc

import (
	"context"

	"github.com/Joe5451/modular-ecommerce/user/internal/application"
	"github.com/Joe5451/modular-ecommerce/user/userpb"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	userpb.UnimplementedUserServiceServer
}

var _ userpb.UserServiceServer = (*server)(nil)

func RegisterServer(app application.App, registrar grpc.ServiceRegistrar) error {
	userpb.RegisterUserServiceServer(registrar, &server{app: app})
	return nil
}

func (s *server) AuthenticateUser(ctx context.Context, req *userpb.AuthenticateUserRequest) (*userpb.AuthenticateUserResponse, error) {
	cmd := application.AuthenticateUser{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	user, err := s.app.AuthenticateUser(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return &userpb.AuthenticateUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	query := application.GetUser{
		ID: req.GetId(),
	}

	user, err := s.app.GetUser(ctx, query)
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *server) RegisterUser(ctx context.Context, req *userpb.RegisterUserRequest) (*userpb.RegisterUserResponse, error) {
	cmd := application.RegisterUser{
		ID:       req.GetId(),
		Email:    req.GetEmail(),
		Name:     req.GetName(),
		Password: req.GetPassword(),
	}

	err := s.app.RegisterUser(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return &userpb.RegisterUserResponse{
		Id:      cmd.ID,
		Message: "User registered successfully",
	}, nil
}
