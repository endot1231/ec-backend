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
	ShopService
}

type UserService interface {
	GetUsers(ctx context.Context, name string, email string) ([]*model.User, error)
	CreateUser(ctx context.Context, name string, email string, password string) (*model.User, error)
}

type ShopService interface {
	GetShops(ctx context.Context, name string, address string) ([]*model.Shop, error)
}

type AuthService interface {
	Login(ctx context.Context, role string, email string, password string) (string, error)
}

type services struct {
	userService
	authService
	shopService
}

func New(exec ent.Client) Services {
	return &services{
		userService: userService{exec: exec},
		authService: authService{exec: exec, Clocker: clock.RealClocker{}},
		shopService: shopService{exec: exec, Clocker: clock.RealClocker{}},
	}
}
