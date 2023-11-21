package redis

import (
	"fmt"
	"webDevScaffold/settings"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.PoolSize,
	})

	// 检验是否ping通
	_, err = rdb.Ping().Result()
	return
}

func Close() (err error) {
	err = rdb.Close()
	if err != nil {
		zap.L().Error("redis close err ! ", zap.Error(err))
		return
	}
	return
}
