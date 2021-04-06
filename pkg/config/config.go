package config

import 	"github.com/thisisdevelopment/xredis/pkg/cache"

type Config struct {
	LogLevel string `yaml:"log"`
	Cache    *cache.Config
}
