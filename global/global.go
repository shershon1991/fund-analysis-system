package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"shershon1991/fund-analye-system/config"
)

// 全局客户端
var (
	GvaConfig      config.ServerConfig // 全局配置
	GvaLogger      *zap.Logger         // 日志
	GvaMysqlClient *gorm.DB            // Mysql客户端
	GvaRedis       *redis.Client       // Redis客户端
	GvaElastic     *elastic.Client     // ES客户端
)
