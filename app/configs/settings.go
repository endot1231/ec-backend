package configs

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Settings struct {
	JwtSecret string `env:"JwtSecret,required"`
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
