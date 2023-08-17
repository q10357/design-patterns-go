package config

import (
	"sync"
)

type config struct {
	cfgs map[string]string
}

var instance *config
var once sync.Once

func GetInstance() *config {
	once.Do(
		func() {
			instance = &config{
				cfgs: make(map[string]string),
			}
		})
	return instance
}

func (c *config) Set(key, value string) {
	instance.cfgs[key] = value
}

func (c *config) Get(key string) string {
	return instance.cfgs[key]
}
