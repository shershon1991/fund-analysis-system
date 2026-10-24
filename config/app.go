// Package config /**
package config

// 应用信息
type app struct {
	ConfigFile string `yaml:"configFile"`
	Desc       string `yaml:"desc"`
	Version    string `yaml:"version"`
	Addr       string `yaml:"addr"`
	Env        string `yaml:"env"`
}

// ServerConfig 配置信息
type ServerConfig struct {
	App     app     `yaml:"app"`
	Mysql   mysql   `yaml:"mysql"`
	Log     log     `yaml:"log"`
	Jwt     jwt     `yaml:"jwt"`
	Redis   redis   `yaml:"redis"`
	Elastic elastic `yaml:"elastic"`
	Cron    cron    `yaml:"cron"`
}
