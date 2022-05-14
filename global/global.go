package global

import (
	"douyin/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Config config.Config
	Db     *gorm.DB
	Rdb    *redis.Client
)
