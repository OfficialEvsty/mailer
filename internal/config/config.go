package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type MailerConfig struct {
	Env      string     `yaml:"env"`
	From     string     `yaml:"from"`
	SmtpHost string     `yaml:"smtp_host"`
	SmtpPort string     `yaml:"smtp_port"`
	GRPC     GRPCConfig `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

// MustLoadPath loads config from yaml file
func MustLoadPath(path string) *MailerConfig {
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config path does not exist: " + path)
	}

	var cfg MailerConfig

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}
