package services

import (
	"context"

	"github.com/endot1231/ec-backend/clock"
	"github.com/endot1231/ec-backend/ent"
	"github.com/endot1231/ec-backend/graph/model"
)

type Services interface {
	UserService
	AuthService
}

type UserService interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUserByName(ctx context.Context, name string) ([]*model.User, error)
	GetUserByAge(ctx context.Context, age int) (*model.User, error)
	CreateUser(ctx context.Context, name string, email string, password string) (*model.User, error)
}

type AuthService interface {
	Login(ctx context.Context, role string, email string, password string) (string, error)
}

type services struct {
	userService
	authService
}

func New(exec ent.Client) Services {
	return &services{
		userService: userService{exec: exec},
		authService: authService{exec: exec, Clocker: clock.RealClocker{}},
	}
}
