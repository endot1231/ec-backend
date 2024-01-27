package services

import (
	"context"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/endot1231/ec-backend/clock"
	"github.com/endot1231/ec-backend/ent"
	"github.com/endot1231/ec-backend/graph/model"
	"github.com/endot1231/ec-backend/test_utils"
	"github.com/stretchr/testify/suite"
)

type TestUsersStruct struct {
	suite.Suite
	db      *ent.Client
	Clocker clock.Clocker

	users      []*ent.Users
	user_role  string
	user_name  string
	user_email string
	user_pass  string

	user1_role  string
	user1_name  string
	user1_email string
	user1_pass  string
}

func TestUsersStructSuite(t *testing.T) {
	os.Setenv("JwtSecret", "12345678")
	suite.Run(t, new(TestUsersStruct))
}

func (s *TestUsersStruct) SetupTest() {
	s.db = test_utils.NewTestDbClient(s.T())
	s.Clocker = clock.FixedClocker{}

	s.user_role = "admin"
	s.user_name = "example"
	s.user_email = "example@example.com"
	s.user_pass = "password1234"

	s.user1_role = "admin"
	s.user1_name = "example1"
	s.user1_email = "example1@example.com"
	s.user1_pass = "password1234"

	ctx := context.Background()
	user1, _ := s.db.Users.Create().SetName(s.user_name).SetEmail(s.user_email).SetPassword(s.user_pass).Save(ctx)
	user2, _ := s.db.Users.Create().SetName(s.user1_name).SetEmail(s.user1_email).SetPassword(s.user1_pass).Save(ctx)
	s.users = append(s.users, user1, user2)
}

func (s *TestUsersStruct) Test_convertUser() {
	want := &model.User{ID: strconv.Itoa(s.users[0].ID), Name: s.users[0].Name, Email: s.users[0].Email}

	type args struct {
		user *ent.Users
	}
	tests := []struct {
		name string
		args args
		want *model.User
	}{
		{
			name: "user is conver",
			args: args{user: s.users[0]},
			want: want,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if got := convertUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *TestUsersStruct) Test_convertUsers() {
	var want []*model.User
	user1 := &model.User{ID: strconv.Itoa(s.users[0].ID), Name: s.users[0].Name, Email: s.users[0].Email}
	user2 := &model.User{ID: strconv.Itoa(s.users[1].ID), Name: s.users[1].Name, Email: s.users[1].Email}
	want = append(want, user1, user2)

	type args struct {
		users []*ent.Users
	}
	tests := []struct {
		name string
		args args
		want []*model.User
	}{
		{
			name: "users is conver",
			args: args{users: s.users},
			want: want,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if got := convertUsers(tt.args.users); !reflect.DeepEqual(got[0], tt.want[0]) && !reflect.DeepEqual(got[1], tt.want[1]) {
				t.Errorf("convertUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *TestUsersStruct) Test_userService_GetUsers() {
	ctx := context.Background()
	var want []*model.User
	user1 := &model.User{ID: strconv.Itoa(s.users[0].ID), Name: s.users[0].Name, Email: s.users[0].Email}
	user2 := &model.User{ID: strconv.Itoa(s.users[1].ID), Name: s.users[1].Name, Email: s.users[1].Email}
	want = append(want, user1, user2)

	type fields struct {
		exec ent.Client
	}
	type args struct {
		ctx   context.Context
		name  string
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.User
		wantErr bool
	}{
		{
			name:    "condition is blank",
			fields:  fields{exec: *s.db},
			args:    args{ctx: ctx, name: "", email: ""},
			want:    want,
			wantErr: false,
		},
		{
			name:    "input name where match both",
			fields:  fields{exec: *s.db},
			args:    args{ctx: ctx, name: "example", email: ""},
			want:    []*model.User{want[0], want[1]},
			wantErr: false,
		},
		{
			name:    "input name where match 1",
			fields:  fields{exec: *s.db},
			args:    args{ctx: ctx, name: "example1", email: ""},
			want:    []*model.User{want[1]},
			wantErr: false,
		},
		{
			name:    "input name where not match both",
			fields:  fields{exec: *s.db},
			args:    args{ctx: ctx, name: "example2", email: ""},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "input email where match both",
			fields:  fields{exec: *s.db},
			args:    args{ctx: ctx, name: "", email: "example"},
			want:    want,
			wantErr: false,
		},
		{
			name:    "input email where match 1",
			fields:  fields{exec: *s.db},
			args:    args{ctx: ctx, name: "", email: "example1"},
			want:    []*model.User{want[1]},
			wantErr: false,
		},
		{
			name:    "input name where not match both",
			fields:  fields{exec: *s.db},
			args:    args{ctx: ctx, name: "", email: "example2"},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "input name and email where match both",
			fields:  fields{exec: *s.db},
			args:    args{ctx: ctx, name: "example", email: "example"},
			want:    want,
			wantErr: false,
		},
		{
			name:    "input name and email where match 1",
			fields:  fields{exec: *s.db},
			args:    args{ctx: ctx, name: "example1", email: "example"},
			want:    []*model.User{want[1]},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			u := &userService{
				exec: tt.fields.exec,
			}
			got, err := u.GetUsers(tt.args.ctx, tt.args.name, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *TestUsersStruct) Test_userService_CreateUser() {
	ctx := context.Background()
	user := &model.User{Name: "name", Email: "email@email.com"}

	type fields struct {
		exec ent.Client
	}
	type args struct {
		ctx      context.Context
		name     string
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name:    "input name and email",
			fields:  fields{exec: *s.db},
			args:    args{ctx: ctx, name: user.Name, email: user.Email, password: "password"},
			want:    user,
			wantErr: false,
		},
		{
			name:    "name is blank",
			fields:  fields{exec: *s.db},
			args:    args{ctx: ctx, name: "", email: user.Email, password: "password"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "email is blank",
			fields:  fields{exec: *s.db},
			args:    args{ctx: ctx, name: user.Name, email: "", password: "password"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "email is not unique",
			fields:  fields{exec: *s.db},
			args:    args{ctx: ctx, name: user.Name, email: s.user_email, password: "password"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "password is blank",
			fields:  fields{exec: *s.db},
			args:    args{ctx: ctx, name: user.Name, email: s.user_email, password: ""},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			u := &userService{
				exec: tt.fields.exec,
			}
			got, err := u.CreateUser(tt.args.ctx, tt.args.name, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				tt.want.ID = got.ID
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("userService.CreateUser() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
