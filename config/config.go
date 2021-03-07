package config

import (
	"os"
	"path/filepath"

	"github.com/valonekowd/clean-architecture/infrastructure/config"
	"github.com/valonekowd/clean-architecture/infrastructure/config/viper"
	"github.com/valonekowd/clean-architecture/util/helper"
)

var AppEnvs = []string{"production", "staging", "testing", "development"}

type Config struct {
	AppEnv string
	Server struct {
		Host string
		Port string
	}
	Datastore struct {
		Primary struct {
			DriverName string
			Host       string
			Port       string
			Username   string
			Password   string
			DBName     string
		}
	}
	Auth struct {
		JWT struct {
			Secret    string
			Issuer    string
			Algorithm string
		}
	}
	Validation struct {
		Playground struct {
			TagName string
		}
	}
}

func (c *Config) IsProd() bool {
	return c.AppEnv == "production"
}

func Create() (*Config, error) {
	configPaths := []string{
		filepath.Join(".", "cmd", "server", "config"),
	}

	c := &Config{}

	var configFiles []*config.File
	{
		f, err := config.NewFile("default", "yml", configPaths)
		if err != nil {
			return nil, err
		}
		configFiles = append(configFiles, f)
	}
	{
		appEnv := os.Getenv("APP_ENV")
		if !helper.StringInSlice(appEnv, AppEnvs) {
			appEnv = "development"
		}

		c.AppEnv = appEnv

		f, err := config.NewFile(c.AppEnv, "yml", configPaths)
		if err != nil {
			return nil, err
		}
		configFiles = append(configFiles, f)
	}

	vc := viper.NewConfiger(
		viper.ConfigerReadFromEnv(true),
		viper.ConfigerFiles(configFiles...),
	)

	return c, vc.Load(c)
}
