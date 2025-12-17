package config

import (
	"github.com/HypercodeHQ/Hypercode/env"
)

type Config struct {
	HTTPAddr string
}

func New() Config {
	return Config{
		HTTPAddr: env.GetVar("HTTP_ADDR", ":8080"),
	}
}
