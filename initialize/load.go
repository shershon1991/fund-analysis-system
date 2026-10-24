package initialize

// 启动加载
func SetLoadInit() {
	// 初始化全局配置文件
	initConfig()
	// 初始化日志zap
	initLogger()
	// 初始化数据库gorm
	initGorm()
	// 初始化缓存redis
	initRedis()
	// 初始化es
	initElastic()
	// 初始化定时任务cron
	initCron()
}
