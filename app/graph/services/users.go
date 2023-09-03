package services

import (
	"context"
	"fmt"
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
		ID:   strconv.Itoa(user.ID),
		Name: user.Name,
	}
}

func convertUsers(users []*ent.Users) []*model.User {
	var userArray []*model.User
	for _, user := range users {
		userArray = append(userArray, convertUser(user))
	}
	return userArray
}

func (u *userService) GetUsers(ctx context.Context) ([]*model.User, error) {
	users, err := u.exec.Users.Query().All(ctx)
	fmt.Println(users)
	if err != nil {
		fmt.Println("stdout")
		return nil, err
	}

	return convertUsers(users), nil
}

func (u *userService) GetUserByName(ctx context.Context, name string) ([]*model.User, error) {
	users, err := u.exec.Users.Query().Where(users.Name(name)).All(ctx)
	fmt.Println(users)
	if err != nil {
		fmt.Println("stdout")
		fmt.Println(err)
		return nil, err
	}

	return convertUsers(users), nil
}

func (u *userService) GetUserByAge(ctx context.Context, age int) (*model.User, error) {
	user, err := u.exec.Users.Query().Where(users.Name("age")).First(ctx)

	if err != nil {
		fmt.Println("stdout")
		fmt.Println(err)
		return nil, err
	}

	return convertUser(user), nil
}

func (u *userService) CreateUser(ctx context.Context, name string, email string, password string) (*model.User, error) {
	user, err := u.exec.Users.
		Create().
		SetName(name).
		SetEmail(email).
		SetPassword(password).
		Save(ctx)

	if err != nil {
		log.Fatalf("failed create company: %v", err)
		return nil, err
	}
	// 3. 戻り値の*model.User型を作る
	return convertUser(user), nil
}
