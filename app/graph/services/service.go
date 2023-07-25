package services

import (
	"context"

	"github.com/endot1231/ec-backend/graph/model"
)

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/$GOPACKAGE/service_mock.go
type Services interface {
	UserService
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	CreateUser(ctx context.Context, name string, email string) (*model.User, error)
}

type services struct {
	*userService
}

func New() Services {
	return &services{
		userService: &userService{},
	}
}
