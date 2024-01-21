package database

import (
	"time"

	"github.com/endot1231/ec-backend/configs"
	"github.com/endot1231/ec-backend/ent"
	"github.com/go-sql-driver/mysql"
)

func NewClient() (*ent.Client, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	c := mysql.Config{
		DBName:    configs.Get().DbName,
		User:      configs.Get().DbUserName,
		Passwd:    configs.Get().DbUserPass,
		Addr:      configs.Get().DbHostName,
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	client, err := ent.Open("mysql", c.FormatDSN())
	return client, err
}
