package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Port string `mapstructure:"PORT"`
}

type Config struct {
	App AppConfig `mapstructure:",squash"`
}

var (
	config  Config
	once    sync.Once
	loadErr error
)

func LoadConfig() (Config, error) {
	once.Do(func() {
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		viper.AddConfigPath("..")
		viper.AddConfigPath("../..")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				loadErr = fmt.Errorf("failed to read config file: %w", err)
				return
			}
		}
		var cfg Config

		if err := viper.Unmarshal(&cfg); err != nil {
			loadErr = fmt.Errorf("failed to unmarshal config: %w", err)
			return
		}
		config = cfg
	})
	return config, loadErr
}
