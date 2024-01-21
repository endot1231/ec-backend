package services

import (
	"context"
	"log"
	"strconv"

	"github.com/endot1231/ec-backend/ent"
	"github.com/endot1231/ec-backend/ent/users"
	"github.com/endot1231/ec-backend/graph/model"
)

type userService struct {
	exec ent.Client
}

func convertUser(user *ent.Users) *model.User {
	return &model.User{
		ID:    strconv.Itoa(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}
}

func convertUsers(users []*ent.Users) []*model.User {
	var userArray []*model.User
	for _, user := range users {
		userArray = append(userArray, convertUser(user))
	}
	return userArray
}

func (u *userService) GetUsers(ctx context.Context, name string, email string) ([]*model.User, error) {
	query := u.exec.Users.Query()

	if name != "" {
		query = query.Where(users.NameEQ(name))
	}

	if email != "" {
		query = query.Where(users.EmailEQ(email))
	}

	users, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	return convertUsers(users), nil
}

func (u *userService) CreateUser(ctx context.Context, name string, email string, password string) (*model.User, error) {
	user, err := u.exec.Users.
		Create().
		SetName(name).
		SetEmail(email).
		SetPassword(password).
		Save(ctx)

	if err != nil {
		log.Fatalf("failed create user: %v", err)
		return nil, err
	}

	return convertUser(user), nil
}
