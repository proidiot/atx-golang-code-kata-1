package main

import "github.com/caarlos0/env"

// Config holds configuration values from the environment, with sane defaults
// (where possible). Required configuration will throw a Fatal error if they
// are missing.
type Config struct {
	AppName string `env:"APP_NAME" envDefault:""`
	AppEnv  string `env:"APP_ENV" envDefault:"development"`
	Port    string `env:"PORT" envDefault:"5000"`
}

// NewConfig returns an instance of Config, with values loaded from ENV vars.
func NewConfig() (Config, error) {
	c := Config{}
	err := env.Parse(&c)
	return c, err
}
