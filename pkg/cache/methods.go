package cache

import (
  "context"
  "time"
  "bytes"
  "github.com/pkg/errors"
)

func (c *Redis) Set(ctx context.Context, key string, value interface{}) error {
  return c.redis.Set(ctx, key, value, time.Duration(c.config.Expiration) * time.Minute).Err()
}

func (c *Redis) Get(ctx context.Context, key string) ([]byte, error) {
  b, err := c.redis.Get(ctx, key).Bytes()
  if err != nil {
    return nil, err
  }

  return b, nil
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

func (c *Redis) Keys2Values(ctx context.Context, keys []string) ([]byte, error) {
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
