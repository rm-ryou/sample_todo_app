package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Port string `mapstructure:"PORT"`
		DB   DB     `mapstructure:",squash"`
	}

	DB struct {
		Database string `mapstructure:"MYSQL_DATABASE"`
		User     string `mapstructure:"MYSQL_USER"`
		Password string `mapstructure:"MYSQL_PASSWORD"`
		Host     string `mapstructure:"MYSQL_HOST"`
		Port     string `mapstructure:"MYSQL_PORT"`
	}
)

func NewConfig() (*Config, error) {
	viper := viper.New()
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	viper.SetDefault("PORT", "8080")
	viper.SetDefault("MYSQL_HOST", "mysql")
	viper.SetDefault("MYSQL_PORT", "3306")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Failed to reading config file: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal to decode into struct: %v", err)
	}

	return &cfg, nil
}
