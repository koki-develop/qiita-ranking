package config

import (
	"os"

	"github.com/caarlos0/env/v8"
	"gopkg.in/yaml.v3"
)

type Config struct {
	QiitaAccessToken string `env:"QIITA_ACCESS_TOKEN,required" yaml:"-"`
	Likes            struct {
		Daily       *ConfigItem `yaml:"daily"`
		DailyByTag  ConfigItems `yaml:"daily_by_tag"`
		Weekly      *ConfigItem `yaml:"weekly"`
		WeeklyByTag ConfigItems `yaml:"weekly_by_tag"`
	} `yaml:"likes"`
}

type ConfigItem struct {
	ItemID string `yaml:"item_id"`
	Tag    string `yaml:"tag"`
}

type ConfigItems []*ConfigItem

func (items ConfigItems) Get(tag string) (*ConfigItem, bool) {
	for _, item := range items {
		if item.Tag == tag {
			return item, true
		}
	}

	return nil, false
}

func Load() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	f, err := os.Open("./config.yml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
