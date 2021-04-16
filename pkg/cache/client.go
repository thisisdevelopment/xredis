package cache

import (
  "context"
  "crypto/tls"
  "time"

  "github.com/go-redis/redis/v8"
  "github.com/logrusorgru/aurora"
  "github.com/thisisdevelopment/go-dockly/xlogger"
)

// ICache defines and exposes the caching layer
type ICache interface {
  Set(ctx context.Context, key string, value interface{}) error
  Get(ctx context.Context, key string, value interface{}) error
  Scan(ctx context.Context, keyname string) ([]string, error)
}

// Redis implements the ICache interface based on redis
type Redis struct {
  redis  *redis.Client
  config *Config
  log    *xlogger.Logger
}

type Config struct {
  Host       string
  Pass       string
  DB         int
  Expiration int
  PoolSize   int `yaml:"pool_size"`
  MaxRetries int `yaml:"max_retries"`
  TLS        bool
}

// New constructs a cache class
func New(config *Config, log *xlogger.Logger) (ICache, error) {

	opts := &redis.Options{
		Addr:       config.Host,
		Password:   config.Pass,
		DB:         config.DB,
		PoolSize:   config.PoolSize,
		MaxRetries: config.MaxRetries,
	}

	if config.TLS == true {
		opts.TLSConfig = &tls.Config{}
	}

	client := redis.NewClient(opts)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	log.Printf("connected to redis %s\n", aurora.Cyan(config.Host))

	return &Redis{
		config: config,
		log:    log,
		redis:  client,
	}, nil
}
