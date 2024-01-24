package main

import (
	"github.com/caarlos0/env/v6"
)

type config struct {
	CoreURL  string `env:"CONF_CORE_URL" envDefault:"http://plinko-core:8080"`
	HTTPPort string `env:"CONF_HTTP_PORT" envDefault:":8080"`
}

func newConfig() (*config, error) {
	conf := new(config)
	err := env.Parse(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
