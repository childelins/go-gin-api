package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	vp *viper.Viper
}

func NewConfig(configFile string) (*Config, error) {
	vp := viper.New()
	vp.SetConfigFile(configFile)
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{vp}, nil
}

func (c *Config) Unmarshal(v interface{}) error {
	return c.vp.Unmarshal(v)
}

