package global

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

var RDB *redis.Client
