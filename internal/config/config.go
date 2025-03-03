package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Server struct {
	Host           string        `mapstructure:"host"`
	Port           int           `mapstructure:"port"`
	Uptime         time.Duration `mapstructure:"uptime"`
	AllowedOrigins string        `mapstructure:"allowed_origins"`
}

type Config struct {
	Server Server `mapstructure:"server"`
}

func NewConfigService() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/app/configs/")
	viper.AddConfigPath("configs/")
	viper.BindEnv("server.host", "RESOURCES_HOST")
	viper.BindEnv("server.port", "RESOURCES_PORT")
	viper.BindEnv("server.uptime", "RESOURCES_UPTIME")
	viper.BindEnv("server.allowed_origins", "RESOURCES_ALLOWED_ORIGINS")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if strings.Contains(err.Error(), "Not Found in") {
			fmt.Println("Config file not found; ignore error if running in CI/CD")
		} else if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found; ignore error if running in CI/CD")
		} else {
			panic(err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	fmt.Println("Config loaded successfully")

	return &cfg
}
