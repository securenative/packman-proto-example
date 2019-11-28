package server

import (
	"github.com/caarlos0/env"
)

type Config struct {
	GrpcPort       int    `env:"GRPC_PORT" envDefault:"{{{ .Port }}}"`
	MonitoringPort int    `env:"MONITORING_PORT" envDefault:"9123"`
}

func ParseConfig() Config {
	out := Config{}

	err := env.Parse(&out)
	if err != nil {
		panic(err)
	}

	return out
}
