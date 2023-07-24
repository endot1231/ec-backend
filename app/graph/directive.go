package graph

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"

	"github.com/endot1231/ec-backend/middlewares/auth"
)

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	if _, ok := auth.GetUserName(ctx); !ok {
		return nil, errors.New("not authenticated")
	}
	return next(ctx)
}
