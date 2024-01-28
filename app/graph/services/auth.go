package services

import (
	"context"
	"errors"
	"strconv"

	"github.com/endot1231/ec-backend/clock"
	"github.com/endot1231/ec-backend/ent"
	"github.com/endot1231/ec-backend/ent/schema"
	"github.com/endot1231/ec-backend/ent/shops"
	"github.com/endot1231/ec-backend/ent/users"
	"github.com/endot1231/ec-backend/secrets"
)

type authService struct {
	exec    ent.Client
	Clocker clock.Clocker
}

func (u *authService) Login(ctx context.Context, role string, email string, password string) (string, error) {

	user, err := u.exec.Users.Query().Where(
		users.EmailEQ(email),
	).Only(ctx)

	if err != nil {
		return "", errors.New(err.Error())
	}

	if user == nil {
		return "", errors.New("failed login")
	}

	err = secrets.CompareValue(user.Password, password)
	if err != nil {
		return "", errors.New(err.Error())
	}

	jwt, err := secrets.JwtGenerate(schema.JwtUsersRole, strconv.Itoa(user.ID), u.Clocker.Now())
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (u *authService) ShopLogin(ctx context.Context, email string, password string) (string, error) {

	shop, err := u.exec.Shops.Query().Where(
		shops.EmailEQ(email),
	).Only(ctx)

	if err != nil {
		return "", errors.New(err.Error())
	}

	if shop == nil {
		return "", errors.New("failed login")
	}

	err = secrets.CompareValue(shop.Password, password)
	if err != nil {
		return "", errors.New(err.Error())
	}

	jwt, err := secrets.JwtGenerate(schema.JwtShopsRole, strconv.Itoa(shop.ID), u.Clocker.Now())
	if err != nil {
		return "", err
	}

	return jwt, nil
}
