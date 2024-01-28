package test_utils

import (
	"testing"

	"github.com/endot1231/ec-backend/ent"
	"github.com/endot1231/ec-backend/ent/enttest"
	"github.com/endot1231/ec-backend/ent/migrate"
)

func NewTestDbClient(t *testing.T) *ent.Client {

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	return client
}
