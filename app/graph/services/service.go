package services

import (
	"context"

	"github.com/endot1231/ec-backend/ent"
	"github.com/endot1231/ec-backend/graph/model"
)

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/$GOPACKAGE/service_mock.go
type Services interface {
	UserService
}

type UserService interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUserByName(ctx context.Context, name string) ([]*model.User, error)
	GetUserByAge(ctx context.Context, age int) (*model.User, error)
	CreateUser(ctx context.Context, name string, email string, password string) (*model.User, error)
}

type services struct {
	userService
}

func New(exec ent.Client) Services {
	return &services{
		userService: userService{exec: exec},
	}
}
