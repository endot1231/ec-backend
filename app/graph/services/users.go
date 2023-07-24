package services

import (
	"context"
	"strconv"

	"github.com/endot1231/ec-backend/graph/db"
	"github.com/endot1231/ec-backend/graph/model"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userService struct {
	exec boil.ContextExecutor
}

func convertUser(user *db.User) *model.User {
	return &model.User{
		ID:   strconv.Itoa(int(user.ID.Int64)),
		Name: user.Name.String,
	}
}

func (u *userService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	// 1. SQLBoilerで生成されたORMコードを呼び出す
	user, err := db.Users( // from users
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name), // select id, name
		db.UserWhere.Name.EQ(null.StringFrom(name)),                 // where name = {引数nameの内容}
	).One(ctx, u.exec) // limit 1
	// 2. エラー処理
	if err != nil {
		return nil, err
	}
	// 3. 戻り値の*model.User型を作る
	return convertUser(user), nil
}

func (u *userService) CreateUser(ctx context.Context, name string, email string) (*model.User, error) {
	// 1. SQLBoilerで生成されたORMコードを呼び出す

	user := db.User{
		Name:  null.StringFrom(name),
		Email: null.StringFrom(email),
	}

	err := user.Insert(ctx, u.exec, boil.Infer())

	// 2. エラー処理
	if err != nil {
		return nil, err
	}
	// 3. 戻り値の*model.User型を作る
	return convertUser(&user), nil
}
