package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/lowl11/planet/log"
	"github.com/lowl11/planet/param"
	"os"
)

type config struct {
	envFiles []string
}

var instance *config

func get(envFiles ...string) *config {
	if instance != nil {
		return instance
	}

	instance = &config{
		envFiles: envFiles,
	}
	instance.load()
	return instance
}

func (config *config) Get(key string) *param.Param {
	return param.New(os.Getenv(key))
}

func (config *config) Parse(result any) error {
	return envconfig.Process("", &result)
}

func (config *config) Load(filesNames ...string) error {
	return godotenv.Load(filesNames...)
}

func (config *config) load() {
	envFileName := ".env"
	if len(config.envFiles) > 0 {
		envFileName = config.envFiles[0]
	}

	_, err := os.Stat(envFileName)
	if os.IsNotExist(err) {
		return
	}

	if err = godotenv.Load(config.envFiles...); err != nil {
		log.Fatal("Load configuration error: ", err)
	}
}
