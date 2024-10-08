package initialize

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"shershon1991/fund-analye-system/global"
)

var configFile string

/**
*  parseCommand
*  @Desc：解析命令行参数
*  @Author Shershon
*  @Date 2021-11-02 10:31:54
**/
func parseCommand() {
	// 读取配置文件优先级: 命令行 > 默认值
	flag.StringVar(&configFile, "c", "./config.yaml", "配置")
	fmt.Println("configFile:", configFile)
	flag.Parse()
}

// ViperInit 初始化viper配置解析包，函数可接受命令行参数
func initConfig() {
	parseCommand()
	if len(configFile) == 0 {
		// 读取默认配置文件
		panic(any("配置文件不存在！"))
	}
	// 读取配置文件
	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		panic(any(fmt.Errorf("配置解析失败:%s\n", err)))
	}
	// 动态监测配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生改变")
		if err := v.Unmarshal(&global.GvaConfig); err != nil {
			panic(any(fmt.Errorf("配置重载失败:%s\n", err)))
		}
	})
	if err := v.Unmarshal(&global.GvaConfig); err != nil {
		panic(any(fmt.Errorf("配置重载失败:%s\n", err)))
	}
	// 设置配置文件
	global.GvaConfig.App.ConfigFile = configFile
}
