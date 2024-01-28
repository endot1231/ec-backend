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
