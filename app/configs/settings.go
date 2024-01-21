package configs

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Settings struct {
	JwtSecret string `env:"JwtSecret,required"`

	DbHostName string `env:"DbHostName" envDefault:"db:3306"`
	DbName     string `env:"DbName" envDefault:"ec_site"`
	DbUserName string `env:"DbUserName" envDefault:"root"`
	DbUserPass string `env:"DbUserPass" envDefault:"root"`

	TestDbHostName string `env:"TestDbHostName" envDefault:"db:3306"`
	TestDbName     string `env:"TestDbName" envDefault:"ec_site_test"`
	TestDbUserName string `env:"TestDbUserName" envDefault:"root"`
	TestDbUserPass string `env:"TestDbUserPass" envDefault:"root"`
}

var cfg Settings = Settings{}

func Init() {
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func Get() Settings {
	return cfg
}
