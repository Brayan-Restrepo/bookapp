package config

import (
	"flag"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"sync"
	"time"
)

type (
	Database struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		DBName   string `mapstructure:"dbname" validate:"required"`
		SSLMode  string `mapstructure:"sslmode" validate:"required"`
		Schema   string `mapstructure:"schema" validate:"required"`
	}

	Server struct {
		Port         string        `mapstructure:"port" validate:"required"`
		AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		Timeout      time.Duration `mapstructure:"timeout" validate:"required"`
		BodyLimit    string        `mapstructure:"bodyLimit" validate:"required"`
		PrefixPath   string        `mapstructure:"prefixPath" validate:"required"`
	}
	Config struct {
		Server   *Server   `mapstructure:"server" validate:"required"`
		Database *Database `mapstructure:"database" validate:"required"`
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func Getting() *Config {
	once.Do(func() {
		env := flag.String("env", "local", "define el entorno (local, dev, pdn)")
		flag.Parse()

		viper.SetConfigName(*env)
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config/env")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}

		validate := validator.New()

		if err := validate.Struct(configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
