package config

import (
	"github.com/lowl11/planet/log"
	"github.com/lowl11/planet/param"
)

func Get(key string) *param.Param {
	return get().Get(key)
}

func Parse(result any) error {
	return get().Parse(result)
}

func Load(filesNames ...string) error {
	return get().Load(filesNames...)
}

func MustParse(result any) {
	if err := get().Parse(result); err != nil {
		log.Fatal("Load config structure error: ", err)
	}
}
