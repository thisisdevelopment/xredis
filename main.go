package main

import (
  "github.com/thisisdevelopment/xredis/pkg/config"
  "github.com/thisisdevelopment/xredis/pkg/cache"
  "github.com/namsral/flag"
  "github.com/thisisdevelopment/go-dockly/xconfig"
  "github.com/thisisdevelopment/go-dockly/xerrors/iferr"
  "github.com/thisisdevelopment/go-dockly/xlogger"
  "fmt"
  "context"
  "time"
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

  ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
  defer cancel()

  err = c.Set(ctx, "redis_1", 1)
  iferr.Exit(err)

  err = c.Set(ctx, "redis_2", 2)
  iferr.Exit(err)

  err = c.Set(ctx, "redis_3", 3)
  iferr.Exit(err)

  g, err := c.Get(ctx, "redis_1")
  iferr.Exit(err)

  s, err := c.Scan(ctx, "redis*")
  iferr.Exit(err)

  v, err := c.Keys2Values(ctx, s)
  iferr.Exit(err)

  fmt.Println(string(g))
  fmt.Println(s)
  fmt.Println(string(v))
}
