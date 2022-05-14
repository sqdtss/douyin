package initialize

import (
	"douyin/global"
	"github.com/go-redis/redis/v8"
)

// Redis 配置Redis
func Redis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.DB, // use default DB
	})
	global.Rdb = rdb
}
