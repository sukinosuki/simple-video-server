package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"strings"
)

type envConfig struct {
	Name  string
	Port  int
	Debug bool
}

type mysqlConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type clickHouseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type redisConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database int
}

type jwtConfig struct {
	SecretKey     string
	EffectiveTime int
}

type logConfig struct {
	MaxAge     int
	MaxSize    int
	MaxBackups int
}

type elasticsearchConfig struct {
	Host string
	Port int
}

type ossConfig struct {
	Protocol        string
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}

type emailConfig struct {
	AuthCode string
	Email    string
}

type appConfig struct {
	Env           envConfig
	Mysql         mysqlConfig
	Jwt           jwtConfig
	Redis         redisConfig
	Oss           ossConfig
	Log           logConfig
	Email         emailConfig
	Clickhouse    clickHouseConfig
	Elasticsearch elasticsearchConfig
}

var (
	Env           envConfig
	Mysql         mysqlConfig
	Jwt           jwtConfig
	Redis         redisConfig
	Oss           ossConfig
	Log           logConfig
	Email         emailConfig
	Clickhouse    clickHouseConfig
	Elasticsearch elasticsearchConfig
)

func init() {

	var _appConfig appConfig

	_, err := toml.DecodeFile("config/config.toml", &_appConfig)

	if err != nil {
		panic(err)
	}

	// 处理命令参数
	args := os.Args

	for _, v := range args {
		arr := strings.Split(v, "=")
		switch arr[0] {
		// 命令参数替换配置值
		case "release":
			if arr[1] == "true" {
				_appConfig.Env.Debug = false
			}
		}
	}

	Env = _appConfig.Env

	Mysql = _appConfig.Mysql

	Jwt = _appConfig.Jwt

	Redis = _appConfig.Redis

	Oss = _appConfig.Oss

	Log = _appConfig.Log

	Email = _appConfig.Email

	Clickhouse = _appConfig.Clickhouse

	Elasticsearch = _appConfig.Elasticsearch
}
