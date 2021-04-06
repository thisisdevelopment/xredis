package main

import (
	"github.com/thisisdevelopment/xredis/pkg/config"
	"github.com/thisisdevelopment/xredis/pkg/cache"
	"github.com/namsral/flag"
	"github.com/thisisdevelopment/go-dockly/xconfig"
	"github.com/thisisdevelopment/go-dockly/xerrors/iferr"
	"github.com/thisisdevelopment/go-dockly/xlogger"
)

var defaultConfig = &xlogger.Config{
	Level:  "debug",
	Format: "text",
}

func main() {

	cfgPath := ""
	flag.StringVar(&cfgPath, "cfg", "config/dev.yaml", "Path to the config file")
	flag.Parse()

	cfg := new(config.Config)
	err := xconfig.LoadConfig(cfg, cfgPath)
	iferr.Exit(err)

	defaultConfig.Level = cfg.LogLevel
	log, err := xlogger.New(defaultConfig)
	iferr.Exit(err)

	c, err := cache.New(cfg.Cache, log)
	iferr.Exit(err)

	_ = c
}
