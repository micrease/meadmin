package redis

import (
	"context"
	"giftcard/library/logger"
	"giftcard/library/structs"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

//基于go-redis v8封装
type RedisClient struct {
	rdb *redis.Client
	ctx context.Context
}

var redisClient = new(RedisClient)

func Connect(addr string) *RedisClient {
	if len(addr) <= 0 {
		addr = "127.0.0.1:6379"
	}
	redisClient.rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	redisClient.ctx = context.Background()

	err := Set("test", 1, time.Minute)
	if err != nil {
		logger.Error("Redis Error", zap.Error(err))
	}

	return redisClient
}

//不过期
func SetForever(key string, value any) error {
	return redisClient.rdb.Set(redisClient.ctx, key, value, 0).Err()
}

//设置过期时间,为了防把过期时间误传,如果不到1秒的,会当成秒处理
func Set(key string, value any, expiration time.Duration) error {
	if expiration < time.Second {
		expiration = expiration * time.Second
	}
	return redisClient.rdb.Set(redisClient.ctx, key, value, expiration).Err()
}

func Lock(key string, value any, expiration time.Duration) bool {
	if expiration < time.Second {
		expiration = expiration * time.Second
	}
	cmd := redisClient.rdb.SetNX(redisClient.ctx, key, value, expiration)
	if cmd.Err() != nil {
		return false
	}

	return cmd.Val()
}

//GET
func Get(key string) *redis.StringCmd {
	return redisClient.rdb.Get(redisClient.ctx, key)
}

//key:string,value:struct
func HMSet(key string, values any, expiration time.Duration) *redis.BoolCmd {
	if expiration < time.Second {
		expiration = expiration * time.Second
	}
	return redisClient.rdb.HMSet(redisClient.ctx, key, values)
}

func HMGet(key string) *redis.StringStringMapCmd {
	return redisClient.rdb.HGetAll(redisClient.ctx, key)
}

func Incr(key string, val float64) *redis.FloatCmd {
	return redisClient.rdb.IncrByFloat(redisClient.ctx, key, val)
}

func Delete(key string) *redis.IntCmd {
	return redisClient.rdb.Del(redisClient.ctx, key)
}

func SetStruct(key string, structVal any, expiration time.Duration) *redis.BoolCmd {
	if expiration < time.Second {
		expiration = expiration * time.Second
	}

	rdb := redisClient.rdb
	mapVal := structs.ToMap(structVal)
	boolCmd := rdb.HMSet(redisClient.ctx, key, mapVal)
	redisClient.rdb.Expire(redisClient.ctx, key, expiration)
	return boolCmd
}

func GetStruct[T any](key string) (*T, error) {
	val := HMGet(key)
	t := new(T)
	//需要redis tag
	err := Scan(val, t, "json")
	return t, err
}

func Scan(cmd *redis.StringStringMapCmd, dest interface{}, tag ...string) error {
	if cmd.Err() != nil {
		return cmd.Err()
	}
	strct, err := structs.Struct(dest, tag[0])
	if err != nil {
		return err
	}

	for k, v := range cmd.Val() {
		_, err := strct.Scan(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
