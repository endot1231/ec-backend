package services

import (
	"context"
	"strconv"
	"testing"

	"github.com/endot1231/ec-backend/clock"
	"github.com/endot1231/ec-backend/ent"
	"github.com/endot1231/ec-backend/secrets"
	"github.com/endot1231/ec-backend/test_utils"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

var (
	user_role  = "admin"
	user_name  = "example"
	user_email = "example@example.com"
	user_pass  = "password1234"
)

var user *ent.Users

type TestAuthStruct struct {
	suite.Suite
	db      *ent.Client
	Clocker clock.Clocker
}

func TestAuthStructSuite(t *testing.T) {
	suite.Run(t, new(TestAuthStruct))
}

func (s *TestAuthStruct) SetupTest() {
	s.db = test_utils.NewTestDbClient(s.T())
	s.Clocker = clock.FixedClocker{}
}

func (s *TestAuthStruct) BeforeTest(_, _ string) {
	ctx := context.Background()
	user, _ = s.db.Users.Create().SetName(user_name).SetEmail(user_email).SetPassword(user_pass).Save(ctx)
}

func (s *TestAuthStruct) Test_authService_Login() {
	ctx := context.Background()
	jwt, _ := secrets.JwtGenerate(user_role, strconv.Itoa(user.ID), s.Clocker.Now())
	type fields struct {
		exec ent.Client
	}
	type args struct {
		ctx      context.Context
		role     string
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Currect user is login",
			fields:  fields{exec: *s.db},
			args:    args{ctx: ctx, role: user_role, email: user_email, password: user_pass},
			want:    jwt,
			wantErr: false,
		},
		{
			name:    "email is wrong",
			fields:  fields{exec: *s.db},
			args:    args{ctx: ctx, role: user_role, email: "wrong@example.com", password: user_pass},
			want:    "",
			wantErr: true,
		},
		{
			name:    "password is wrong",
			fields:  fields{exec: *s.db},
			args:    args{ctx: ctx, role: user_role, email: user_email, password: "hogehoge"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			u := &authService{
				exec:    *s.db,
				Clocker: s.Clocker,
			}
			got, err := u.Login(tt.args.ctx, tt.args.role, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("authService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("authService.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
