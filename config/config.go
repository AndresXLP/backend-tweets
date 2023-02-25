package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerHost string `required:"true" split_words:"true"`
	ServerPort int    `required:"true" split_words:"true"`
	DbHost     string `required:"true" split_words:"true"`
	DbPort     int    `required:"true" split_words:"true"`
	DbUser     string `required:"true" split_words:"true"`
	DbPassword string `required:"true" split_words:"true"`
	DbName     string `required:"true" split_words:"true"`
	SecretJWT  string `required:"true" split_words:"true"`
}

var (
	Once sync.Once
	Cfg  Config
)

func Environments() Config {
	Once.Do(func() {
		if err := envconfig.Process("", &Cfg); err != nil {
			log.Panicf("Error parsing environment vars %#v", err)
		}
	})

	return Cfg
}
