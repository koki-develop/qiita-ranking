package cmd

import "github.com/caarlos0/env/v8"

type config struct {
	QiitaAccessToken string `env:"QIITA_ACCESS_TOKEN,required"`
}

func loadConfig() (*config, error) {
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
