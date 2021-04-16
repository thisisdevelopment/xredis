package cache

import (
  "context"
  "time"
  "bytes"
  "fmt"
  "github.com/pkg/errors"
)

func (c *Redis) Set(ctx context.Context, key string, value interface{}) error {
  return c.redis.Set(ctx, key, value, time.Duration(c.config.Expiration) * time.Minute).Err()
}

func (c *Redis) Get(ctx context.Context, key string, value interface{}) error {
  var err error

  switch p := value.(type) {
    case *int:
      *p , err = c.redis.Get(ctx, key).Int()
    case *int64:
      *p, err = c.redis.Get(ctx, key).Int64()
    case *uint64:
      *p, err = c.redis.Get(ctx, key).Uint64()
    case *float32:
      *p, err = c.redis.Get(ctx, key).Float32()
    case *float64:
      *p, err = c.redis.Get(ctx, key).Float64()
    case *bool:
      *p, err = c.redis.Get(ctx, key).Bool()
    case *time.Time:
      *p, err = c.redis.Get(ctx, key).Time()
    case *string:
      *p, err = c.redis.Get(ctx, key).Result()
    default:
      value, err = nil, fmt.Errorf("Dynamic Type not supported\n")
  }

  return err
}

func (c *Redis) Scan(ctx context.Context, keyname string) ([]string, error) {
  var res, keys []string
  var cursor uint64
  var err error

  for {
    res, cursor, err = c.redis.Scan(ctx, cursor, keyname, 10).Result()
    if err != nil {
      return nil, errors.Wrap(err, "failed Redis Scan")
    }

    keys = append(keys, res...)

    if cursor == 0 {
      break
    }
  }

  return keys, nil
}

func (c *Redis) keys2Values(ctx context.Context, keys []string) ([]byte, error) {
  buf := new(bytes.Buffer)
  buf.WriteRune('[')

  for _, key := range keys {
    val, err := c.redis.Get(ctx, key).Bytes()
    if err != nil {
      return nil, errors.Wrap(err, "failed Redis Get")
    }
    buf.Write(val)
    buf.WriteRune(',')
  }

  len := buf.Len()
  if len > 1 {
    buf.Truncate(len - 1)
  }

  buf.WriteRune(']')
  return buf.Bytes(), nil
}
