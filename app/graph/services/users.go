package services

import (
	"context"

	"github.com/endot1231/ec-backend/graph/model"
)

type userService struct {
}

func (u *userService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	// 1. SQLBoilerで生成されたORMコードを呼び出す

	// 3. 戻り値の*model.User型を作る
	return nil, nil
}

func (u *userService) CreateUser(ctx context.Context, name string, email string) (*model.User, error) {
	// 1. SQLBoilerで生成されたORMコードを呼び出す

	// 3. 戻り値の*model.User型を作る
	return nil, nil
}
