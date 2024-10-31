package config

import (
	"os"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Name         string `yaml:"name" envconfig:"DAZED_APP_NAME"`
		LogLevel     string `yaml:"log_level" envconfig:"DAZED_APP_LOG_LEVEL"`
		RootPassword string `yaml:"root_password" envconfig:"DAZEDAPP_ROOT_PASSWORD"`
	} `yaml:"app"`
	Server struct {
		Addr string `yaml:"addr" envconfig:"DAZED_SERVER_ADDR"`
	} `yaml:"server"`
	Postgres struct {
		URL      string `yaml:"url" envconfig:"DAZED_POSTGRES_URL"`
		User     string `yaml:"user" envconfig:"DAZED_POSTGRES_USER"`
		Password string `yaml:"password" envconfig:"DAZED_POSTGRES_PASSWORD"`
		Db       string `yaml:"db" envconfig:"DAZED_POSTGRES_DB"`
		SSLMode  string `yaml:"ssl_mode" envconfig:"DAZED_POSTGRES_SSL_MODE"`
	} `yaml:"postgres"`
	Redis struct {
		URL      string `yaml:"url" envconfig:"DAZED_REDIS_URL"`
		DB       int    `yaml:"db" envconfig:"DAZED_REDIS_DB"`
		Password string `yaml:"password" envconfig:"DAZED_REDIS_PASSWORD"`
	} `yaml:"redis"`
}

var (
	cfg  Config
	once sync.Once
)

func Get() Config {
	once.Do(func() {
		if err := parse(
			parseYAML("config.yaml"),
			parseEnv,
		); err != nil {
			panic(err)
		}
	})
	return cfg
}

func parse(parsers ...parser) error {
	for _, p := range parsers {
		if err := p(&cfg); err != nil {
			return err
		}
	}
	return nil
}

type parser func(*Config) error

func parseYAML(filename string) parser {
	return func(config *Config) error {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		return yaml.NewDecoder(file).Decode(config)
	}
}

func parseEnv(config *Config) error {
	return envconfig.Process("DAZED_", config)
}
