package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	MaxConfig    Max    `mapstructure:"max_config"`
	LoggerConfig Logger `mapstructure:"logger"`
}

type Max struct {
	Token string `mapstructure:"token"`
}

type Logger struct {
	Path  string `mapstructure:"path"`
	Level string `mapstructure:"level"`
}

const (
	ConfigName = "config"
	ConfigType = "yaml"
	ConfigPath = "./config"
)

func LoadConfig() (*Config, error) {
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigType)
	viper.AddConfigPath(ConfigPath)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	fmt.Println("Config file used:", viper.ConfigFileUsed())
	return &cfg, nil
}
