package cache

import (
	"context"
	"time"
)

func (c *Redis) Set(ctx context.Context, key string, value interface{}) error {

	return c.redis.Set(ctx, key, value, time.Duration(c.config.Expiration)*time.Minute).Err()
}

func (c *Redis) Get(ctx context.Context, key string) ([]byte, error) {

	val, err := c.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return []byte(val), nil
}

func (c *Redis) Scan(ctx context.Context, partialKey string) (res []byte, err error) {
	var keys []string
	var cursor uint64 ///!!!important todo handle cursor and limit
	keys, cursor, err = c.redis.Scan(ctx, cursor, partialKey, 100).Result()
	if err != nil {
		return nil, err
	}

	c.log.Dump(partialKey, keys)

	res = append(res, []byte("[")...)

	for _, key := range keys {
		val, err := c.redis.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}
		res = append(res, []byte(val)...)
		res = append(res, []byte(",")...)
	}

	if len(res) > 1 {
		// remove comma from last element and close array
		res = res[:len(res)-1]
	}
	res = append(res, []byte("]")...)

	return res, nil
}
