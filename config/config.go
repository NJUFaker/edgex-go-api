package config

import (
	"fmt"
	"log"

	"github.com/go-ini/ini"
)

var AppConfig struct {
	LogSetting
	MysqlSetting
	MongoSetting
	RedisSetting
	ServerSetting
}

type ServerSetting struct {
	RunMode  string
	HttpPort int
	Port     string
}

type LogSetting struct {
	LogLevel   string
	FileName   string // 日志文件名
	MaxSize    int    // 每个日志文件保存的最大尺寸 单位：M
	MaxBackups int    // 日志文件最多保存多少个备份
	MaxAge     int    // 文件最多保存多少天
	Compress   bool   // 日志是否压缩
}

type MysqlSetting struct {
}

type MongoSetting struct {
}

type RedisSetting struct {
}

func InitAppConfig() {

	cfg, err := ini.Load("app.ini")
	if err != nil {
		panic(err)
	}

	mapTo("Log", &AppConfig.LogSetting, cfg)
	mapTo("Mongo", &AppConfig.MongoSetting, cfg)
	mapTo("Mysql", &AppConfig.MysqlSetting, cfg)
	mapTo("Redis", &AppConfig.RedisSetting, cfg)
	mapTo("Server", &AppConfig.ServerSetting, cfg)

	if AppConfig.HttpPort != 0 {
		AppConfig.Port = fmt.Sprintf(":%d", AppConfig.HttpPort)
	}
}

func mapTo(section string, v interface{}, cfg *ini.File) {
	if cfg == nil || section == "" {
		log.Fatalf("section=%v, iniFile=%v", section, cfg)
		return
	}
	if err := cfg.Section(section).MapTo(v); err != nil {
		panic(err)
	}
}
