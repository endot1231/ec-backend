package services

import (
	"context"

	"github.com/endot1231/ec-backend/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/$GOPACKAGE/service_mock.go
type Services interface {
	UserService
	RepoService
}

type UserService interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type RepoService interface {
	GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error)
}

type services struct {
	*userService
	*repoService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService: &userService{exec: exec},
		repoService: &repoService{exec: exec},
	}
}
