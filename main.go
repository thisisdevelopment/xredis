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

  var i int
  var i64 int64
  var ui64 uint64
  var b bool
  var f32 float32
  var f64 float64
  var s string

  err = c.Set(ctx, "int", 111)
  iferr.Exit(err)
  err = c.Set(ctx, "int64", 222)
  iferr.Exit(err)
  err = c.Set(ctx, "uint64", 333)
  iferr.Exit(err)
  err = c.Set(ctx, "bool", true)
  iferr.Exit(err)
  err = c.Set(ctx, "float32", 444.444)
  iferr.Exit(err)
  err = c.Set(ctx, "float64", 555.555)
  iferr.Exit(err)
  err = c.Set(ctx, "string", "hello world")
  iferr.Exit(err)

  err = c.Get(ctx, "int", &i)
  iferr.Exit(err)
  err = c.Get(ctx, "int64", &i64)
  iferr.Exit(err)
  err = c.Get(ctx, "uint64", &ui64)
  iferr.Exit(err)
  err = c.Get(ctx, "bool", &b)
  iferr.Exit(err)
  err = c.Get(ctx, "float32", &f32)
  iferr.Exit(err)
  err = c.Get(ctx, "float64", &f64)
  iferr.Exit(err)
  err = c.Get(ctx, "string", &s)
  iferr.Exit(err)
  
  fmt.Printf("%T %[1]v\n",i)
  fmt.Printf("%T %[1]v\n",i64)
  fmt.Printf("%T %[1]v\n",ui64)
  fmt.Printf("%T %[1]v\n",b)
  fmt.Printf("%T %[1]v\n",f32)
  fmt.Printf("%T %[1]v\n",f64)
  fmt.Printf("%T %[1]v\n",s)

/*
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
*/
}
